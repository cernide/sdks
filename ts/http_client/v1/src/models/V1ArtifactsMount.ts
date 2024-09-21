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
/**
 *
 * @export
 * @interface V1ArtifactsMount
 */
export interface V1ArtifactsMount {
    /**
     *
     * @type {string}
     * @memberof V1ArtifactsMount
     */
    name?: string;
    /**
     *
     * @type {Array<string>}
     * @memberof V1ArtifactsMount
     */
    paths?: Array<string>;
}

/**
 * Check if a given object implements the V1ArtifactsMount interface.
 */
export function instanceOfV1ArtifactsMount(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1ArtifactsMountFromJSON(json: any): V1ArtifactsMount {
    return V1ArtifactsMountFromJSONTyped(json, false);
}

export function V1ArtifactsMountFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): V1ArtifactsMount {
    if (json === undefined || json === null) {
        return json;
    }
    return {
        name: !exists(json, "name") ? undefined : json["name"],
        paths: !exists(json, "paths") ? undefined : json["paths"],
    };
}

export function V1ArtifactsMountToJSON(value?: V1ArtifactsMount | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        name: value.name,
        paths: value.paths,
    };
}
