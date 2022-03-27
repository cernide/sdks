// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.17.1
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1GridSearch model module.
 * @module model/V1GridSearch
 * @version 1.17.1
 */
class V1GridSearch {
    /**
     * Constructs a new <code>V1GridSearch</code>.
     * @alias module:model/V1GridSearch
     */
    constructor() { 
        
        V1GridSearch.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1GridSearch</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1GridSearch} obj Optional instance to populate.
     * @return {module:model/V1GridSearch} The populated <code>V1GridSearch</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1GridSearch();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('params')) {
                obj['params'] = ApiClient.convertToType(data['params'], {'String': Object});
            }
            if (data.hasOwnProperty('numRuns')) {
                obj['numRuns'] = ApiClient.convertToType(data['numRuns'], 'Number');
            }
            if (data.hasOwnProperty('seed')) {
                obj['seed'] = ApiClient.convertToType(data['seed'], 'Number');
            }
            if (data.hasOwnProperty('concurrency')) {
                obj['concurrency'] = ApiClient.convertToType(data['concurrency'], 'Number');
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
 * @default 'grid'
 */
V1GridSearch.prototype['kind'] = 'grid';

/**
 * @member {Object.<String, Object>} params
 */
V1GridSearch.prototype['params'] = undefined;

/**
 * @member {Number} numRuns
 */
V1GridSearch.prototype['numRuns'] = undefined;

/**
 * @member {Number} seed
 */
V1GridSearch.prototype['seed'] = undefined;

/**
 * @member {Number} concurrency
 */
V1GridSearch.prototype['concurrency'] = undefined;

/**
 * @member {Array.<Object>} earlyStopping
 */
V1GridSearch.prototype['earlyStopping'] = undefined;






export default V1GridSearch;

