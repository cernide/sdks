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

/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.21.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1OptimizationMetric from './V1OptimizationMetric';
import V1Tuner from './V1Tuner';

/**
 * The V1Bayes model module.
 * @module model/V1Bayes
 * @version 1.21.0
 */
class V1Bayes {
    /**
     * Constructs a new <code>V1Bayes</code>.
     * @alias module:model/V1Bayes
     */
    constructor() {

        V1Bayes.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Bayes</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Bayes} obj Optional instance to populate.
     * @return {module:model/V1Bayes} The populated <code>V1Bayes</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Bayes();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('params')) {
                obj['params'] = ApiClient.convertToType(data['params'], {'String': Object});
            }
            if (data.hasOwnProperty('numInitialRuns')) {
                obj['numInitialRuns'] = ApiClient.convertToType(data['numInitialRuns'], 'Number');
            }
            if (data.hasOwnProperty('maxIterations')) {
                obj['maxIterations'] = ApiClient.convertToType(data['maxIterations'], 'Number');
            }
            if (data.hasOwnProperty('utilityFunction')) {
                obj['utilityFunction'] = ApiClient.convertToType(data['utilityFunction'], Object);
            }
            if (data.hasOwnProperty('metric')) {
                obj['metric'] = V1OptimizationMetric.constructFromObject(data['metric']);
            }
            if (data.hasOwnProperty('seed')) {
                obj['seed'] = ApiClient.convertToType(data['seed'], 'Number');
            }
            if (data.hasOwnProperty('concurrency')) {
                obj['concurrency'] = ApiClient.convertToType(data['concurrency'], 'Number');
            }
            if (data.hasOwnProperty('tuner')) {
                obj['tuner'] = V1Tuner.constructFromObject(data['tuner']);
            }
            if (data.hasOwnProperty('earlyStopping')) {
                obj['earlyStopping'] = ApiClient.convertToType(data['earlyStopping'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} kind
 * @default 'bayes'
 */
V1Bayes.prototype['kind'] = 'bayes';

/**
 * @member {Object.<String, Object>} params
 */
V1Bayes.prototype['params'] = undefined;

/**
 * @member {Number} numInitialRuns
 */
V1Bayes.prototype['numInitialRuns'] = undefined;

/**
 * @member {Number} maxIterations
 */
V1Bayes.prototype['maxIterations'] = undefined;

/**
 * @member {Object} utilityFunction
 */
V1Bayes.prototype['utilityFunction'] = undefined;

/**
 * @member {module:model/V1OptimizationMetric} metric
 */
V1Bayes.prototype['metric'] = undefined;

/**
 * @member {Number} seed
 */
V1Bayes.prototype['seed'] = undefined;

/**
 * @member {Number} concurrency
 */
V1Bayes.prototype['concurrency'] = undefined;

/**
 * @member {module:model/V1Tuner} tuner
 */
V1Bayes.prototype['tuner'] = undefined;

/**
 * @member {Array.<Object>} earlyStopping
 */
V1Bayes.prototype['earlyStopping'] = undefined;






export default V1Bayes;

