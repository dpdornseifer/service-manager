/*
 * Copyright 2018 The Service Manager Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package base

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/tidwall/sjson"

	"github.com/gofrs/uuid"

	"github.com/Peripli/service-manager/storage"

	"github.com/Peripli/service-manager/pkg/extension"
	"github.com/Peripli/service-manager/pkg/log"
	"github.com/Peripli/service-manager/pkg/query"
	"github.com/Peripli/service-manager/pkg/types"
	"github.com/Peripli/service-manager/pkg/util"
	"github.com/Peripli/service-manager/pkg/web"
)

const PathParamID = "id"

var _ extension.Interceptable = &Controller{}

type Controller struct {
	resourceBaseURL string
	objectType      types.ObjectType
	repository      storage.Repository
	objectBlueprint func() types.Object

	wrappedCreateInterceptorProvider extension.CreateInterceptorWrapper
	CreateInterceptorProviders       []extension.CreateInterceptorProvider

	wrappedUpdateInterceptorProvider extension.UpdateInterceptorWrapper
	UpdateInterceptorProviders       []extension.UpdateInterceptorProvider

	wrappedDeleteInterceptorProvider extension.DeleteInterceptorWrapper
	DeleteInterceptorProviders       []extension.DeleteInterceptorProvider
}

func (c *Controller) InterceptsType() types.ObjectType {
	return c.objectType
}

func (c *Controller) findCreateProviderPosition(providerName string) int {
	index := -1
	for i, provider := range c.CreateInterceptorProviders {
		if providerName == provider.Name() {
			index = i
			break
		}
	}
	return index
}

func (c *Controller) AddCreateInterceptorProvidersBefore(providerName string, providers ...extension.CreateInterceptorProvider) {
	for _, provider := range providers {
		index := c.findCreateProviderPosition(providerName)
		if index == -1 {
			c.CreateInterceptorProviders = append(c.CreateInterceptorProviders, provider)
		} else {
			c.CreateInterceptorProviders = append(c.CreateInterceptorProviders, nil)
			copy(c.CreateInterceptorProviders[index+1:], c.CreateInterceptorProviders[index:])
			c.CreateInterceptorProviders[index] = provider
		}
	}
	c.wrappedCreateInterceptorProvider = extension.UnionCreateInterceptor(c.CreateInterceptorProviders)
}

func (c *Controller) AddCreateInterceptorProvidersAfter(providerName string, providers ...extension.CreateInterceptorProvider) {
	for _, provider := range providers {
		index := c.findCreateProviderPosition(providerName)
		if index == -1 {
			c.CreateInterceptorProviders = append(c.CreateInterceptorProviders, provider)
		} else {
			index = index + 1
			c.CreateInterceptorProviders = append(c.CreateInterceptorProviders, nil)
			copy(c.CreateInterceptorProviders[index+1:], c.CreateInterceptorProviders[index:])
			c.CreateInterceptorProviders[index] = provider
		}
	}
	c.wrappedCreateInterceptorProvider = extension.UnionCreateInterceptor(c.CreateInterceptorProviders)
}

func (c *Controller) AddCreateInterceptorProviders(providers ...extension.CreateInterceptorProvider) {
	// if c.CreateInterceptorProvider == nil {
	// 	c.CreateInterceptorProvider = extension.UnionCreateInterceptor(providers)
	// } else {
	// 	c.CreateInterceptorProvider = extension.UnionCreateInterceptor(append(providers, c.CreateInterceptorProvider))
	// }
}

func (c *Controller) findUpdateProviderPosition(providerName string) int {
	index := -1
	for i, provider := range c.UpdateInterceptorProviders {
		if providerName == provider.Name() {
			index = i
			break
		}
	}
	return index
}

func (c *Controller) AddUpdateInterceptorProvidersBefore(providerName string, providers ...extension.UpdateInterceptorProvider) {
	for _, provider := range providers {
		index := c.findUpdateProviderPosition(providerName)
		if index == -1 {
			c.UpdateInterceptorProviders = append(c.UpdateInterceptorProviders, provider)
		} else {
			c.UpdateInterceptorProviders = append(c.UpdateInterceptorProviders, nil)
			copy(c.UpdateInterceptorProviders[index+1:], c.UpdateInterceptorProviders[index:])
			c.UpdateInterceptorProviders[index] = provider
		}
	}
	c.wrappedUpdateInterceptorProvider = extension.UnionUpdateInterceptor(c.UpdateInterceptorProviders)
}

func (c *Controller) AddUpdateInterceptorProvidersAfter(providerName string, providers ...extension.UpdateInterceptorProvider) {
	for _, provider := range providers {
		index := c.findUpdateProviderPosition(providerName)
		if index == -1 {
			c.UpdateInterceptorProviders = append(c.UpdateInterceptorProviders, provider)
		} else {
			index = index + 1
			c.UpdateInterceptorProviders = append(c.UpdateInterceptorProviders, nil)
			copy(c.UpdateInterceptorProviders[index+1:], c.UpdateInterceptorProviders[index:])
			c.UpdateInterceptorProviders[index] = provider
		}
	}
	c.wrappedUpdateInterceptorProvider = extension.UnionUpdateInterceptor(c.UpdateInterceptorProviders)
}

func (c *Controller) AddUpdateInterceptorProviders(providers ...extension.UpdateInterceptorProvider) {
	// 	if c.UpdateInterceptorProvider == nil {
	// 		c.UpdateInterceptorProvider = extension.UnionUpdateInterceptor(providers)
	// 	} else {
	// 		c.UpdateInterceptorProvider = extension.UnionUpdateInterceptor(append(providers, c.UpdateInterceptorProvider))
	// 	}
}

func (c *Controller) findDeleteProviderPosition(providerName string) int {
	index := -1
	for i, provider := range c.DeleteInterceptorProviders {
		if providerName == provider.Name() {
			index = i
			break
		}
	}
	return index
}

func (c *Controller) AddDeleteInterceptorProvidersBefore(providerName string, providers ...extension.DeleteInterceptorProvider) {
	for _, provider := range providers {
		index := c.findDeleteProviderPosition(providerName)
		if index == -1 {
			c.DeleteInterceptorProviders = append(c.DeleteInterceptorProviders, provider)
		} else {
			c.DeleteInterceptorProviders = append(c.DeleteInterceptorProviders, nil)
			copy(c.DeleteInterceptorProviders[index+1:], c.DeleteInterceptorProviders[index:])
			c.DeleteInterceptorProviders[index] = provider
		}
	}
	c.wrappedDeleteInterceptorProvider = extension.UnionDeleteInterceptor(c.DeleteInterceptorProviders)
}

func (c *Controller) AddDeleteInterceptorProviders(providers ...extension.DeleteInterceptorProvider) {
	// 	if c.DeleteInterceptorProvider == nil {
	// 		c.DeleteInterceptorProvider = extension.UnionDeleteInterceptor(providers)
	// 	} else {
	// 		c.DeleteInterceptorProvider = extension.UnionDeleteInterceptor(append(providers, c.DeleteInterceptorProvider))
	// 	}
}

func NewController(repository storage.Repository, resourceBaseURL string, objectBlueprint func() types.Object) *Controller {
	return &Controller{
		repository:      repository,
		resourceBaseURL: resourceBaseURL,
		objectBlueprint: objectBlueprint,
		objectType:      objectBlueprint().GetType(),
	}
}

func (c *Controller) Routes() []web.Route {
	return []web.Route{
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPost,
				Path:   c.resourceBaseURL,
			},
			Handler: c.CreateObject,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   fmt.Sprintf("%s/{%s}", c.resourceBaseURL, PathParamID),
			},
			Handler: c.GetSingleObject,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodGet,
				Path:   c.resourceBaseURL,
			},
			Handler: c.ListObjects,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path:   c.resourceBaseURL,
			},
			Handler: c.DeleteObjects,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodDelete,
				Path:   fmt.Sprintf("%s/{%s}", c.resourceBaseURL, PathParamID),
			},
			Handler: c.DeleteSingleObject,
		},
		{
			Endpoint: web.Endpoint{
				Method: http.MethodPatch,
				Path:   fmt.Sprintf("%s/{%s}", c.resourceBaseURL, PathParamID),
			},
			Handler: c.PatchObject,
		},
	}
}

func (c *Controller) CreateObject(r *web.Request) (*web.Response, error) {
	ctx := r.Context()
	log.C(ctx).Debugf("Creating new %s", c.objectType)

	var createInterceptor extension.CreateInterceptor
	// TODO: Lazy initialize wrapped providers on first request
	if c.wrappedCreateInterceptorProvider != nil {
		createInterceptor = c.wrappedCreateInterceptorProvider()
	}

	onTransaction := func(ctx context.Context, txStorage storage.Warehouse, newObject types.Object) error {
		id, err := txStorage.Create(ctx, newObject)
		if err != nil {
			return util.HandleStorageError(err, string(c.objectType))
		}
		newObject.SetID(id)
		return nil
	}
	if createInterceptor != nil {
		onTransaction = createInterceptor.OnTransactionCreate(onTransaction)
	}

	onAPI := func(ctx context.Context, obj types.Object) (types.Object, error) {
		if err := c.repository.InTransaction(ctx, func(ctx context.Context, txStorage storage.Warehouse) error {
			return onTransaction(ctx, txStorage, obj)
		}); err != nil {
			return nil, err
		}
		return obj, nil
	}
	if createInterceptor != nil {
		onAPI = createInterceptor.OnAPICreate(onAPI)
	}
	result := c.objectBlueprint()
	if err := util.BytesToObject(r.Body, result); err != nil {
		return nil, err
	}

	if result.GetID() == "" {
		UUID, err := uuid.NewV4()
		if err != nil {
			return nil, fmt.Errorf("could not generate GUID for %s: %s", c.objectType, err)
		}
		result.SetID(UUID.String())
	}
	currentTime := time.Now().UTC()
	result.SetCreatedAt(currentTime)
	result.SetUpdatedAt(currentTime)

	result, err := onAPI(ctx, result)
	if err != nil {
		return nil, err
	}
	return web.NewJSONResponse(http.StatusCreated, result)
}

func (c *Controller) DeleteObjects(r *web.Request) (*web.Response, error) {
	ctx := r.Context()
	log.C(ctx).Debugf("Deleting %ss...", c.objectType)
	var deleteInterceptor extension.DeleteInterceptor
	if c.wrappedDeleteInterceptorProvider != nil {
		deleteInterceptor = c.wrappedDeleteInterceptorProvider()
	}

	transactionOperation := func(ctx context.Context, txStorage storage.Warehouse, deletionCriteria ...query.Criterion) (types.ObjectList, error) {
		return c.repository.Delete(ctx, c.objectType, deletionCriteria...)
	}
	if deleteInterceptor != nil {
		transactionOperation = deleteInterceptor.OnTransactionDelete(transactionOperation)
	}

	apiOperation := func(ctx context.Context, deletionCriteria ...query.Criterion) (types.ObjectList, error) {
		var result types.ObjectList
		if err := c.repository.InTransaction(ctx, func(ctx context.Context, storage storage.Warehouse) error {
			var err error
			result, err = transactionOperation(ctx, storage, deletionCriteria...)
			return util.HandleSelectionError(err, string(c.objectType))
		}); err != nil {
			return nil, err
		}
		return result, nil
	}
	if deleteInterceptor != nil {
		apiOperation = deleteInterceptor.OnAPIDelete(apiOperation)
	}
	criteria := query.CriteriaForContext(ctx)
	if _, err := apiOperation(ctx, criteria...); err != nil {
		return nil, err
	}

	return web.NewJSONResponse(http.StatusOK, map[string]string{})
}

func (c *Controller) DeleteSingleObject(r *web.Request) (*web.Response, error) {
	objectID := r.PathParams[PathParamID]
	ctx := r.Context()
	log.C(ctx).Debugf("Deleting %s with id %s", c.objectType, objectID)

	byID := query.ByField(query.EqualsOperator, "id", objectID)
	ctx, err := query.AddCriteria(r.Context(), byID)
	if err != nil {
		return nil, err
	}
	r.Request = r.WithContext(ctx)
	return c.DeleteObjects(r)
}

func (c *Controller) GetSingleObject(r *web.Request) (*web.Response, error) {
	objectID := r.PathParams[PathParamID]
	ctx := r.Context()
	log.C(ctx).Debugf("Getting %s with id %s", c.objectType, objectID)

	object, err := c.repository.Get(ctx, c.objectType, objectID)
	if err != nil {
		return nil, util.HandleStorageError(err, string(c.objectType))
	}
	if secured, ok := object.(types.Secured); ok {
		secured.SetCredentials(nil)
	} else {
		log.C(ctx).Debugf("Object of type %s with id %s is not secured, so no credentials are cleaned up on response", object.GetType(), object.GetID())
	}
	return web.NewJSONResponse(http.StatusOK, object)
}

func (c *Controller) ListObjects(r *web.Request) (*web.Response, error) {
	ctx := r.Context()
	log.C(ctx).Debugf("Getting all %ss", c.objectType)
	objectList, err := c.repository.List(ctx, c.objectType, query.CriteriaForContext(ctx)...)
	if err != nil {
		return nil, util.HandleSelectionError(err)
	}
	for i := 0; i < objectList.Len(); i++ {
		obj := objectList.ItemAt(i)
		if secured, ok := obj.(types.Secured); ok {
			secured.SetCredentials(nil)
		} else {
			log.C(ctx).Debugf("Object of type %s with id %s is not secured, so no credentials are cleaned up on response", obj.GetType(), obj.GetID())
		}
	}

	return web.NewJSONResponse(http.StatusOK, objectList)
}

func (c *Controller) PatchObject(r *web.Request) (*web.Response, error) {
	objectID := r.PathParams[PathParamID]
	ctx := r.Context()
	log.C(ctx).Debugf("Updating %s with id %s", c.objectType, objectID)

	labelChanges, err := query.LabelChangesFromJSON(r.Body)
	if err != nil {
		return nil, err
	}
	if r.Body, err = sjson.DeleteBytes(r.Body, "labels"); err != nil {
		return nil, err
	}

	var updateInterceptor extension.UpdateInterceptor
	if c.wrappedUpdateInterceptorProvider != nil {
		updateInterceptor = c.wrappedUpdateInterceptorProvider()
	}
	transactionOp := func(ctx context.Context, txStorage storage.Warehouse, oldObject types.Object, updateChanges *extension.UpdateContext) (types.Object, error) {
		return txStorage.Update(ctx, oldObject, updateChanges.LabelChanges...)
	}
	if updateInterceptor != nil {
		transactionOp = updateInterceptor.OnTransactionUpdate(transactionOp)
	}

	objFromDB, err := c.repository.Get(ctx, c.objectType, objectID)
	if err != nil {
		return nil, util.HandleStorageError(err, string(c.objectType))
	}
	createdAt := objFromDB.GetCreatedAt()

	apiOperation := func(ctx context.Context, updateChanges *extension.UpdateContext) (types.Object, error) {
		var result types.Object
		if err := util.BytesToObject(updateChanges.ObjectChanges, updateChanges.Object); err != nil {
			return nil, err
		}
		updateChanges.Object.SetID(objectID)
		updateChanges.Object.SetCreatedAt(createdAt)
		updateChanges.Object.SetUpdatedAt(time.Now().UTC())
		if err = c.repository.InTransaction(ctx, func(ctx context.Context, txStorage storage.Warehouse) error {
			result, err = transactionOp(ctx, txStorage, updateChanges.Object, updateChanges)
			return util.HandleStorageError(err, string(c.objectType))
		}); err != nil {
			return nil, err
		}
		return result, nil
	}
	if updateInterceptor != nil {
		apiOperation = updateInterceptor.OnAPIUpdate(apiOperation)
	}

	objectChanges := &extension.UpdateContext{
		LabelChanges:  labelChanges,
		ObjectChanges: r.Body,
		Object:        objFromDB,
	}
	object, err := apiOperation(ctx, objectChanges)
	if err != nil {
		return nil, err
	}
	if obj, ok := object.(types.Secured); ok {
		obj.SetCredentials(nil)
	} else {
		log.C(ctx).Debugf("Object of type %s with id %s is not secured, so no credentials are cleaned up on response", object.GetType(), object.GetID())
	}
	return web.NewJSONResponse(http.StatusOK, object)
}