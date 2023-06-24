/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc22
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
export const V1TriggerPolicy = {
    AllSucceeded: 'all_succeeded',
    AllFailed: 'all_failed',
    AllDone: 'all_done',
    OneSucceeded: 'one_succeeded',
    OneFailed: 'one_failed',
    OneDone: 'one_done'
} as const;
export type V1TriggerPolicy = typeof V1TriggerPolicy[keyof typeof V1TriggerPolicy];


export function V1TriggerPolicyFromJSON(json: any): V1TriggerPolicy {
    return V1TriggerPolicyFromJSONTyped(json, false);
}

export function V1TriggerPolicyFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1TriggerPolicy {
    return json as V1TriggerPolicy;
}

export function V1TriggerPolicyToJSON(value?: V1TriggerPolicy | null): any {
    return value as any;
}

