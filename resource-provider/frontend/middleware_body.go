package frontend

// Copyright (c) Microsoft Corporation.
// Licensed under the Apache License 2.0.

import (
	"context"
	"github.com/Azure/ARO-HCP/resource-provider/api/arm"
	"io"
	"net/http"
	"strings"
)

func MiddlewareBody(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	switch r.Method {
	case http.MethodPatch, http.MethodPost, http.MethodPut:
		body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 1048576))
		if err != nil {
			arm.WriteError(
				w, http.StatusBadRequest,
				arm.CloudErrorCodeInvalidResource, "",
				"The resource definition is invalid.")
			return
		}

		contentType := strings.SplitN(r.Header.Get("Content-Type"), ";", 2)[0]

		if contentType != "application/json" && !(len(body) == 0 && contentType == "") {
			arm.WriteError(
				w, http.StatusUnsupportedMediaType,
				arm.CloudErrorCodeUnsupportedMediaType, "",
				"The content media type '%s' is not supported. Only 'application/json' is supported.",
				r.Header.Get("Content-Type"))
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), ContextKeyBody, body))
	}

	next(w, r)
}
