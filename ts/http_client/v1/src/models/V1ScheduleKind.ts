/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc28
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
export const V1ScheduleKind = {
    Cron: 'cron',
    Interval: 'interval',
    Datetime: 'datetime'
} as const;
export type V1ScheduleKind = typeof V1ScheduleKind[keyof typeof V1ScheduleKind];


export function V1ScheduleKindFromJSON(json: any): V1ScheduleKind {
    return V1ScheduleKindFromJSONTyped(json, false);
}

export function V1ScheduleKindFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ScheduleKind {
    return json as V1ScheduleKind;
}

export function V1ScheduleKindToJSON(value?: V1ScheduleKind | null): any {
    return value as any;
}

