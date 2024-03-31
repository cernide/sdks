/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.4
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 *
 * @export
 */
export const V1CleanPodPolicy = {
    All: 'All',
    Running: 'Running',
    None: 'None'
} as const;
export type V1CleanPodPolicy = typeof V1CleanPodPolicy[keyof typeof V1CleanPodPolicy];


export function V1CleanPodPolicyFromJSON(json: any): V1CleanPodPolicy {
    return V1CleanPodPolicyFromJSONTyped(json, false);
}

export function V1CleanPodPolicyFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1CleanPodPolicy {
    return json as V1CleanPodPolicy;
}

export function V1CleanPodPolicyToJSON(value?: V1CleanPodPolicy | null): any {
    return value as any;
}

