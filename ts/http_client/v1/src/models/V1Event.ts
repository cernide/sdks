/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import type { V1EventArtifact } from './V1EventArtifact';
import {
    V1EventArtifactFromJSON,
    V1EventArtifactFromJSONTyped,
    V1EventArtifactToJSON,
} from './V1EventArtifact';
import type { V1EventAudio } from './V1EventAudio';
import {
    V1EventAudioFromJSON,
    V1EventAudioFromJSONTyped,
    V1EventAudioToJSON,
} from './V1EventAudio';
import type { V1EventChart } from './V1EventChart';
import {
    V1EventChartFromJSON,
    V1EventChartFromJSONTyped,
    V1EventChartToJSON,
} from './V1EventChart';
import type { V1EventConfusionMatrix } from './V1EventConfusionMatrix';
import {
    V1EventConfusionMatrixFromJSON,
    V1EventConfusionMatrixFromJSONTyped,
    V1EventConfusionMatrixToJSON,
} from './V1EventConfusionMatrix';
import type { V1EventCurve } from './V1EventCurve';
import {
    V1EventCurveFromJSON,
    V1EventCurveFromJSONTyped,
    V1EventCurveToJSON,
} from './V1EventCurve';
import type { V1EventDataframe } from './V1EventDataframe';
import {
    V1EventDataframeFromJSON,
    V1EventDataframeFromJSONTyped,
    V1EventDataframeToJSON,
} from './V1EventDataframe';
import type { V1EventHistogram } from './V1EventHistogram';
import {
    V1EventHistogramFromJSON,
    V1EventHistogramFromJSONTyped,
    V1EventHistogramToJSON,
} from './V1EventHistogram';
import type { V1EventImage } from './V1EventImage';
import {
    V1EventImageFromJSON,
    V1EventImageFromJSONTyped,
    V1EventImageToJSON,
} from './V1EventImage';
import type { V1EventModel } from './V1EventModel';
import {
    V1EventModelFromJSON,
    V1EventModelFromJSONTyped,
    V1EventModelToJSON,
} from './V1EventModel';
import type { V1EventSpan } from './V1EventSpan';
import {
    V1EventSpanFromJSON,
    V1EventSpanFromJSONTyped,
    V1EventSpanToJSON,
} from './V1EventSpan';
import type { V1EventVideo } from './V1EventVideo';
import {
    V1EventVideoFromJSON,
    V1EventVideoFromJSONTyped,
    V1EventVideoToJSON,
} from './V1EventVideo';

/**
 *
 * @export
 * @interface V1Event
 */
export interface V1Event {
    /**
     *
     * @type {Date}
     * @memberof V1Event
     */
    timestamp?: Date;
    /**
     * Global step of the event.
     * @type {number}
     * @memberof V1Event
     */
    step?: number;
    /**
     *
     * @type {number}
     * @memberof V1Event
     */
    metric?: number;
    /**
     *
     * @type {V1EventImage}
     * @memberof V1Event
     */
    image?: V1EventImage;
    /**
     *
     * @type {V1EventHistogram}
     * @memberof V1Event
     */
    histogram?: V1EventHistogram;
    /**
     *
     * @type {V1EventAudio}
     * @memberof V1Event
     */
    audio?: V1EventAudio;
    /**
     *
     * @type {V1EventVideo}
     * @memberof V1Event
     */
    video?: V1EventVideo;
    /**
     *
     * @type {string}
     * @memberof V1Event
     */
    html?: string;
    /**
     *
     * @type {string}
     * @memberof V1Event
     */
    text?: string;
    /**
     *
     * @type {V1EventChart}
     * @memberof V1Event
     */
    chart?: V1EventChart;
    /**
     *
     * @type {V1EventModel}
     * @memberof V1Event
     */
    model?: V1EventModel;
    /**
     *
     * @type {V1EventArtifact}
     * @memberof V1Event
     */
    artifact?: V1EventArtifact;
    /**
     *
     * @type {V1EventDataframe}
     * @memberof V1Event
     */
    dataframe?: V1EventDataframe;
    /**
     *
     * @type {V1EventCurve}
     * @memberof V1Event
     */
    curve?: V1EventCurve;
    /**
     *
     * @type {V1EventConfusionMatrix}
     * @memberof V1Event
     */
    confusion?: V1EventConfusionMatrix;
    /**
     *
     * @type {V1EventSpan}
     * @memberof V1Event
     */
    span?: V1EventSpan;
}

/**
 * Check if a given object implements the V1Event interface.
 */
export function instanceOfV1Event(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function V1EventFromJSON(json: any): V1Event {
    return V1EventFromJSONTyped(json, false);
}

export function V1EventFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Event {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {

        'timestamp': !exists(json, 'timestamp') ? undefined : (new Date(json['timestamp'])),
        'step': !exists(json, 'step') ? undefined : json['step'],
        'metric': !exists(json, 'metric') ? undefined : json['metric'],
        'image': !exists(json, 'image') ? undefined : V1EventImageFromJSON(json['image']),
        'histogram': !exists(json, 'histogram') ? undefined : V1EventHistogramFromJSON(json['histogram']),
        'audio': !exists(json, 'audio') ? undefined : V1EventAudioFromJSON(json['audio']),
        'video': !exists(json, 'video') ? undefined : V1EventVideoFromJSON(json['video']),
        'html': !exists(json, 'html') ? undefined : json['html'],
        'text': !exists(json, 'text') ? undefined : json['text'],
        'chart': !exists(json, 'chart') ? undefined : V1EventChartFromJSON(json['chart']),
        'model': !exists(json, 'model') ? undefined : V1EventModelFromJSON(json['model']),
        'artifact': !exists(json, 'artifact') ? undefined : V1EventArtifactFromJSON(json['artifact']),
        'dataframe': !exists(json, 'dataframe') ? undefined : V1EventDataframeFromJSON(json['dataframe']),
        'curve': !exists(json, 'curve') ? undefined : V1EventCurveFromJSON(json['curve']),
        'confusion': !exists(json, 'confusion') ? undefined : V1EventConfusionMatrixFromJSON(json['confusion']),
        'span': !exists(json, 'span') ? undefined : V1EventSpanFromJSON(json['span']),
    };
}

export function V1EventToJSON(value?: V1Event | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {

        'timestamp': value.timestamp === undefined ? undefined : (value.timestamp.toISOString()),
        'step': value.step,
        'metric': value.metric,
        'image': V1EventImageToJSON(value.image),
        'histogram': V1EventHistogramToJSON(value.histogram),
        'audio': V1EventAudioToJSON(value.audio),
        'video': V1EventVideoToJSON(value.video),
        'html': value.html,
        'text': value.text,
        'chart': V1EventChartToJSON(value.chart),
        'model': V1EventModelToJSON(value.model),
        'artifact': V1EventArtifactToJSON(value.artifact),
        'dataframe': V1EventDataframeToJSON(value.dataframe),
        'curve': V1EventCurveToJSON(value.curve),
        'confusion': V1EventConfusionMatrixToJSON(value.confusion),
        'span': V1EventSpanToJSON(value.span),
    };
}

