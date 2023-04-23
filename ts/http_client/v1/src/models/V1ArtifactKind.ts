// Copyright 2018-2023 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 2.0.0-rc12
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * - model: Model asset/event
 *  - audio: Audio asset/event
 *  - video: Vidio asset/event
 *  - histogram: Histogram asset/event
 *  - image: Image asset/event
 *  - tensor: Tensor asset/event
 *  - dataframe: Dataframe asset/event
 *  - chart: plotly/bokeh/vega chart
 *  - csv: Comma separated values
 *  - tsv: Tab separated values
 *  - psv: Pipe separated values
 *  - ssv: Space separated values
 *  - metric: Metric asset/event
 *  - env: Env file
 *  - html: HTML asset/event
 *  - text: Text asset/event
 *  - file: File asset/lineage
 *  - dir: Dir asset/lineage
 *  - dockerfile: Dockerfile asset
 *  - docker_image: Docker image
 *  - data: Data asset/event
 *  - coderef: Coderef lineage
 *  - table: Table asset/event
 *  - tensorboard: Tensorboard lineage
 *  - curve: Curve event
 *  - confusion: Confusion matrix event
 *  - analysis: Analysis lineage
 *  - iteration: Iteration lineage
 *  - markdown: Mardown event
 *  - system: System event
 *  - artifact: Generic artifact
 * @export
 */
export const V1ArtifactKind = {
    Model: 'model',
    Audio: 'audio',
    Video: 'video',
    Histogram: 'histogram',
    Image: 'image',
    Tensor: 'tensor',
    Dataframe: 'dataframe',
    Chart: 'chart',
    Csv: 'csv',
    Tsv: 'tsv',
    Psv: 'psv',
    Ssv: 'ssv',
    Metric: 'metric',
    Env: 'env',
    Html: 'html',
    Text: 'text',
    File: 'file',
    Dir: 'dir',
    Dockerfile: 'dockerfile',
    DockerImage: 'docker_image',
    Data: 'data',
    Coderef: 'coderef',
    Table: 'table',
    Tensorboard: 'tensorboard',
    Curve: 'curve',
    Confusion: 'confusion',
    Analysis: 'analysis',
    Iteration: 'iteration',
    Markdown: 'markdown',
    System: 'system',
    Artifact: 'artifact'
} as const;
export type V1ArtifactKind = typeof V1ArtifactKind[keyof typeof V1ArtifactKind];


export function V1ArtifactKindFromJSON(json: any): V1ArtifactKind {
    return V1ArtifactKindFromJSONTyped(json, false);
}

export function V1ArtifactKindFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1ArtifactKind {
    return json as V1ArtifactKind;
}

export function V1ArtifactKindToJSON(value?: V1ArtifactKind | null): any {
    return value as any;
}

