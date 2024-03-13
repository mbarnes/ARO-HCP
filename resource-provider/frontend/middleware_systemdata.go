package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"encoding/json"
	"github.com/Azure/ARO-HCP/resource-provider/api/arm"
	"net/http"
)

func MiddlewareSystemData(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// See https://eng.ms/docs/products/arm/api_contracts/resourcesystemdata
	data := r.Header.Get("X-Ms-Arm-Resource-System-Data")
	if data != "" {
		var systemData arm.SystemData
		err := json.Unmarshal([]byte(data), &systemData)
		if err != nil {
			// FIXME Log the error.
		}
		r = r.WithContext(context.WithValue(r.Context(), ContextKeySystemData, systemData))
	}

	next(w, r)
}
