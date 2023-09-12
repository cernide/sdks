/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc41
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1ArtifactsMount } from './V1ArtifactsMount';
import {
    V1ArtifactsMountFromJSON,
    V1ArtifactsMountFromJSONTyped,
    V1ArtifactsMountToJSON,
} from './V1ArtifactsMount';
import type { V1ArtifactsType } from './V1ArtifactsType';
import {
    V1ArtifactsTypeFromJSON,
    V1ArtifactsTypeFromJSONTyped,
    V1ArtifactsTypeToJSON,
} from './V1ArtifactsType';
import type { V1AuthType } from './V1AuthType';
import {
    V1AuthTypeFromJSON,
    V1AuthTypeFromJSONTyped,
    V1AuthTypeToJSON,
} from './V1AuthType';
import type { V1CompiledOperation } from './V1CompiledOperation';
import {
    V1CompiledOperationFromJSON,
    V1CompiledOperationFromJSONTyped,
    V1CompiledOperationToJSON,
} from './V1CompiledOperation';
import type { V1ConnectionResource } from './V1ConnectionResource';
import {
    V1ConnectionResourceFromJSON,
    V1ConnectionResourceFromJSONTyped,
    V1ConnectionResourceToJSON,
} from './V1ConnectionResource';
import type { V1ConnectionSchema } from './V1ConnectionSchema';
import {
    V1ConnectionSchemaFromJSON,
    V1ConnectionSchemaFromJSONTyped,
    V1ConnectionSchemaToJSON,
} from './V1ConnectionSchema';
import type { V1ConnectionType } from './V1ConnectionType';
import {
    V1ConnectionTypeFromJSON,
    V1ConnectionTypeFromJSONTyped,
    V1ConnectionTypeToJSON,
} from './V1ConnectionType';
import type { V1EarlyStopping } from './V1EarlyStopping';
import {
    V1EarlyStoppingFromJSON,
    V1EarlyStoppingFromJSONTyped,
    V1EarlyStoppingToJSON,
} from './V1EarlyStopping';
import type { V1Event } from './V1Event';
import {
    V1EventFromJSON,
    V1EventFromJSONTyped,
    V1EventToJSON,
} from './V1Event';
import type { V1EventType } from './V1EventType';
import {
    V1EventTypeFromJSON,
    V1EventTypeFromJSONTyped,
    V1EventTypeToJSON,
} from './V1EventType';
import type { V1GcsType } from './V1GcsType';
import {
    V1GcsTypeFromJSON,
    V1GcsTypeFromJSONTyped,
    V1GcsTypeToJSON,
} from './V1GcsType';
import type { V1HpParams } from './V1HpParams';
import {
    V1HpParamsFromJSON,
    V1HpParamsFromJSONTyped,
    V1HpParamsToJSON,
} from './V1HpParams';
import type { V1Matrix } from './V1Matrix';
import {
    V1MatrixFromJSON,
    V1MatrixFromJSONTyped,
    V1MatrixToJSON,
} from './V1Matrix';
import type { V1MatrixKind } from './V1MatrixKind';
import {
    V1MatrixKindFromJSON,
    V1MatrixKindFromJSONTyped,
    V1MatrixKindToJSON,
} from './V1MatrixKind';
import type { V1Operation } from './V1Operation';
import {
    V1OperationFromJSON,
    V1OperationFromJSONTyped,
    V1OperationToJSON,
} from './V1Operation';
import type { V1PolyaxonInitContainer } from './V1PolyaxonInitContainer';
import {
    V1PolyaxonInitContainerFromJSON,
    V1PolyaxonInitContainerFromJSONTyped,
    V1PolyaxonInitContainerToJSON,
} from './V1PolyaxonInitContainer';
import type { V1PolyaxonSidecarContainer } from './V1PolyaxonSidecarContainer';
import {
    V1PolyaxonSidecarContainerFromJSON,
    V1PolyaxonSidecarContainerFromJSONTyped,
    V1PolyaxonSidecarContainerToJSON,
} from './V1PolyaxonSidecarContainer';
import type { V1Reference } from './V1Reference';
import {
    V1ReferenceFromJSON,
    V1ReferenceFromJSONTyped,
    V1ReferenceToJSON,
} from './V1Reference';
import type { V1RunSchema } from './V1RunSchema';
import {
    V1RunSchemaFromJSON,
    V1RunSchemaFromJSONTyped,
    V1RunSchemaToJSON,
} from './V1RunSchema';
import type { V1S3Type } from './V1S3Type';
import {
    V1S3TypeFromJSON,
    V1S3TypeFromJSONTyped,
    V1S3TypeToJSON,
} from './V1S3Type';
import type { V1Schedule } from './V1Schedule';
import {
    V1ScheduleFromJSON,
    V1ScheduleFromJSONTyped,
    V1ScheduleToJSON,
} from './V1Schedule';
import type { V1ScheduleKind } from './V1ScheduleKind';
import {
    V1ScheduleKindFromJSON,
    V1ScheduleKindFromJSONTyped,
    V1ScheduleKindToJSON,
} from './V1ScheduleKind';
import type { V1UriType } from './V1UriType';
import {
    V1UriTypeFromJSON,
    V1UriTypeFromJSONTyped,
    V1UriTypeToJSON,
} from './V1UriType';
import type { V1WasbType } from './V1WasbType';
import {
    V1WasbTypeFromJSON,
    V1WasbTypeFromJSONTyped,
    V1WasbTypeToJSON,
} from './V1WasbType';

/**
 *
 * @export
 * @interface V1Schemas
 */
export interface V1Schemas {
    /**
     *
     * @type {V1EarlyStopping}
     * @memberof V1Schemas
     */
    earlyStopping?: V1EarlyStopping;
    /**
     *
     * @type {V1Matrix}
     * @memberof V1Schemas
     */
    matrix?: V1Matrix;
    /**
     *
     * @type {V1RunSchema}
     * @memberof V1Schemas
     */
    run?: V1RunSchema;
    /**
     *
     * @type {V1Operation}
     * @memberof V1Schemas
     */
    operation?: V1Operation;
    /**
     *
     * @type {V1CompiledOperation}
     * @memberof V1Schemas
     */
    compiledOperation?: V1CompiledOperation;
    /**
     *
     * @type {V1Schedule}
     * @memberof V1Schemas
     */
    schedule?: V1Schedule;
    /**
     *
     * @type {V1ConnectionSchema}
     * @memberof V1Schemas
     */
    connectionSchema?: V1ConnectionSchema;
    /**
     *
     * @type {V1HpParams}
     * @memberof V1Schemas
     */
    hpParams?: V1HpParams;
    /**
     *
     * @type {V1Reference}
     * @memberof V1Schemas
     */
    reference?: V1Reference;
    /**
     *
     * @type {V1ArtifactsMount}
     * @memberof V1Schemas
     */
    artifactsMount?: V1ArtifactsMount;
    /**
     *
     * @type {V1PolyaxonSidecarContainer}
     * @memberof V1Schemas
     */
    polyaxonSidecarContainer?: V1PolyaxonSidecarContainer;
    /**
     *
     * @type {V1PolyaxonInitContainer}
     * @memberof V1Schemas
     */
    polyaxonInitContainer?: V1PolyaxonInitContainer;
    /**
     *
     * @type {V1ArtifactsType}
     * @memberof V1Schemas
     */
    artifacs?: V1ArtifactsType;
    /**
     *
     * @type {V1WasbType}
     * @memberof V1Schemas
     */
    wasb?: V1WasbType;
    /**
     *
     * @type {V1GcsType}
     * @memberof V1Schemas
     */
    gcs?: V1GcsType;
    /**
     *
     * @type {V1S3Type}
     * @memberof V1Schemas
     */
    s3?: V1S3Type;
    /**
     *
     * @type {V1AuthType}
     * @memberof V1Schemas
     */
    auth?: V1AuthType;
    /**
     *
     * @type {V1UriType}
     * @memberof V1Schemas
     */
    uri?: V1UriType;
    /**
     *
     * @type {V1ConnectionResource}
     * @memberof V1Schemas
     */
    resource?: V1ConnectionResource;
    /**
     *
     * @type {V1ConnectionType}
     * @memberof V1Schemas
     */
    connection?: V1ConnectionType;
    /**
     *
     * @type {V1EventType}
     * @memberof V1Schemas
     */
    eventType?: V1EventType;
    /**
     *
     * @type {V1MatrixKind}
     * @memberof V1Schemas
     */
    matrixKind?: V1MatrixKind;
    /**
     *
     * @type {V1ScheduleKind}
     * @memberof V1Schemas
     */
    scheduleKind?: V1ScheduleKind;
    /**
     *
     * @type {V1Event}
     * @memberof V1Schemas
     */
    event?: V1Event;
}

/**
 * Check if a given object implements the V1Schemas interface.
 */
export function instanceOfV1Schemas(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1SchemasFromJSON(json: any): V1Schemas {
    return V1SchemasFromJSONTyped(json, false);
}

export function V1SchemasFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Schemas {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'earlyStopping': !exists(json, 'earlyStopping') ? undefined : V1EarlyStoppingFromJSON(json['earlyStopping']),
        'matrix': !exists(json, 'matrix') ? undefined : V1MatrixFromJSON(json['matrix']),
        'run': !exists(json, 'run') ? undefined : V1RunSchemaFromJSON(json['run']),
        'operation': !exists(json, 'operation') ? undefined : V1OperationFromJSON(json['operation']),
        'compiledOperation': !exists(json, 'compiledOperation') ? undefined : V1CompiledOperationFromJSON(json['compiledOperation']),
        'schedule': !exists(json, 'schedule') ? undefined : V1ScheduleFromJSON(json['schedule']),
        'connectionSchema': !exists(json, 'connectionSchema') ? undefined : V1ConnectionSchemaFromJSON(json['connectionSchema']),
        'hpParams': !exists(json, 'hpParams') ? undefined : V1HpParamsFromJSON(json['hpParams']),
        'reference': !exists(json, 'reference') ? undefined : V1ReferenceFromJSON(json['reference']),
        'artifactsMount': !exists(json, 'artifactsMount') ? undefined : V1ArtifactsMountFromJSON(json['artifactsMount']),
        'polyaxonSidecarContainer': !exists(json, 'polyaxonSidecarContainer') ? undefined : V1PolyaxonSidecarContainerFromJSON(json['polyaxonSidecarContainer']),
        'polyaxonInitContainer': !exists(json, 'polyaxonInitContainer') ? undefined : V1PolyaxonInitContainerFromJSON(json['polyaxonInitContainer']),
        'artifacs': !exists(json, 'artifacs') ? undefined : V1ArtifactsTypeFromJSON(json['artifacs']),
        'wasb': !exists(json, 'wasb') ? undefined : V1WasbTypeFromJSON(json['wasb']),
        'gcs': !exists(json, 'gcs') ? undefined : V1GcsTypeFromJSON(json['gcs']),
        's3': !exists(json, 's3') ? undefined : V1S3TypeFromJSON(json['s3']),
        'auth': !exists(json, 'auth') ? undefined : V1AuthTypeFromJSON(json['auth']),
        'uri': !exists(json, 'uri') ? undefined : V1UriTypeFromJSON(json['uri']),
        'resource': !exists(json, 'resource') ? undefined : V1ConnectionResourceFromJSON(json['resource']),
        'connection': !exists(json, 'connection') ? undefined : V1ConnectionTypeFromJSON(json['connection']),
        'eventType': !exists(json, 'eventType') ? undefined : V1EventTypeFromJSON(json['eventType']),
        'matrixKind': !exists(json, 'matrixKind') ? undefined : V1MatrixKindFromJSON(json['matrixKind']),
        'scheduleKind': !exists(json, 'scheduleKind') ? undefined : V1ScheduleKindFromJSON(json['scheduleKind']),
        'event': !exists(json, 'event') ? undefined : V1EventFromJSON(json['event']),
    };
}

export function V1SchemasToJSON(value?: V1Schemas | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'earlyStopping': V1EarlyStoppingToJSON(value.earlyStopping),
        'matrix': V1MatrixToJSON(value.matrix),
        'run': V1RunSchemaToJSON(value.run),
        'operation': V1OperationToJSON(value.operation),
        'compiledOperation': V1CompiledOperationToJSON(value.compiledOperation),
        'schedule': V1ScheduleToJSON(value.schedule),
        'connectionSchema': V1ConnectionSchemaToJSON(value.connectionSchema),
        'hpParams': V1HpParamsToJSON(value.hpParams),
        'reference': V1ReferenceToJSON(value.reference),
        'artifactsMount': V1ArtifactsMountToJSON(value.artifactsMount),
        'polyaxonSidecarContainer': V1PolyaxonSidecarContainerToJSON(value.polyaxonSidecarContainer),
        'polyaxonInitContainer': V1PolyaxonInitContainerToJSON(value.polyaxonInitContainer),
        'artifacs': V1ArtifactsTypeToJSON(value.artifacs),
        'wasb': V1WasbTypeToJSON(value.wasb),
        'gcs': V1GcsTypeToJSON(value.gcs),
        's3': V1S3TypeToJSON(value.s3),
        'auth': V1AuthTypeToJSON(value.auth),
        'uri': V1UriTypeToJSON(value.uri),
        'resource': V1ConnectionResourceToJSON(value.resource),
        'connection': V1ConnectionTypeToJSON(value.connection),
        'eventType': V1EventTypeToJSON(value.eventType),
        'matrixKind': V1MatrixKindToJSON(value.matrixKind),
        'scheduleKind': V1ScheduleKindToJSON(value.scheduleKind),
        'event': V1EventToJSON(value.event),
    };
}

