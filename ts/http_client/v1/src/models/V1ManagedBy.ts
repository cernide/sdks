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

/**
 *
 * @export
 */
export const V1ManagedBy = {
    Client: "client",
    Cli: "cli",
    Agent: "agent",
} as const;
export type V1ManagedBy = (typeof V1ManagedBy)[keyof typeof V1ManagedBy];

export function V1ManagedByFromJSON(json: any): V1ManagedBy {
    return V1ManagedByFromJSONTyped(json, false);
}

export function V1ManagedByFromJSONTyped(
    json: any,
    ignoreDiscriminator: boolean
): V1ManagedBy {
    return json as V1ManagedBy;
}

export function V1ManagedByToJSON(value?: V1ManagedBy | null): any {
    return value as any;
}
