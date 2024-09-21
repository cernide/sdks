/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.4.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from "../runtime";
import type { V1RunEdge } from "./V1RunEdge";
import {
    V1RunEdgeFromJSON,
    V1RunEdgeFromJSONTyped,
    V1RunEdgeToJSON,
} from "./V1RunEdge";

/**
 *
 * @export
 * @interface V1ListRunEdgesResponse
 */
export interface V1ListRunEdgesResponse {
    /**
     *
     * @type {number}
     * @memberof V1ListRunEdgesResponse
     */
    count?: number;
    /**
     *
     * @type {Array<V1RunEdge>}
     * @memberof V1ListRunEdgesResponse
     */
    results?: Array<V1RunEdge>;
    /**
     *
     * @type {string}
     * @memberof V1ListRunEdgesResponse
     */
    previous?: string;
    /**
     *
     * @type {string}
     * @memberof V1ListRunEdgesResponse
     */
    next?: string;
}

/**
 * Check if a given object implements the V1ListRunEdgesResponse interface.
 */
export function instanceOfV1ListRunEdgesResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ListRunEdgesResponseFromJSON(
    json: any
): V1ListRunEdgesResponse {
    return V1ListRunEdgesResponseFromJSONTyped(json, false);
}

export function V1ListRunEdgesResponseFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): V1ListRunEdgesResponse {
    if (json === undefined || json === null) {
        return json;
    }
    return {
        count: !exists(json, "count") ? undefined : json["count"],
        results: !exists(json, "results")
            ? undefined
            : (json["results"] as Array<any>).map(V1RunEdgeFromJSON),
        previous: !exists(json, "previous") ? undefined : json["previous"],
        next: !exists(json, "next") ? undefined : json["next"],
    };
}

export function V1ListRunEdgesResponseToJSON(
    value?: V1ListRunEdgesResponse | null
): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        count: value.count,
        results:
            value.results === undefined
                ? undefined
                : (value.results as Array<any>).map(V1RunEdgeToJSON),
        previous: value.previous,
        next: value.next,
    };
}
