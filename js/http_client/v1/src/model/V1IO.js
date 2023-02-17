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
 * The version of the OpenAPI document: 1.22.0-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1IO model module.
 * @module model/V1IO
 * @version 1.22.0-rc0
 */
class V1IO {
    /**
     * Constructs a new <code>V1IO</code>.
     * @alias module:model/V1IO
     */
    constructor() { 
        
        V1IO.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1IO</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1IO} obj Optional instance to populate.
     * @return {module:model/V1IO} The populated <code>V1IO</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1IO();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('type')) {
                obj['type'] = ApiClient.convertToType(data['type'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], Object);
            }
            if (data.hasOwnProperty('isOptional')) {
                obj['isOptional'] = ApiClient.convertToType(data['isOptional'], 'Boolean');
            }
            if (data.hasOwnProperty('isList')) {
                obj['isList'] = ApiClient.convertToType(data['isList'], 'Boolean');
            }
            if (data.hasOwnProperty('isFlag')) {
                obj['isFlag'] = ApiClient.convertToType(data['isFlag'], 'Boolean');
            }
            if (data.hasOwnProperty('argFormat')) {
                obj['argFormat'] = ApiClient.convertToType(data['argFormat'], 'String');
            }
            if (data.hasOwnProperty('delayValidation')) {
                obj['delayValidation'] = ApiClient.convertToType(data['delayValidation'], 'Boolean');
            }
            if (data.hasOwnProperty('options')) {
                obj['options'] = ApiClient.convertToType(data['options'], [Object]);
            }
            if (data.hasOwnProperty('connection')) {
                obj['connection'] = ApiClient.convertToType(data['connection'], 'String');
            }
            if (data.hasOwnProperty('toInit')) {
                obj['toInit'] = ApiClient.convertToType(data['toInit'], 'Boolean');
            }
            if (data.hasOwnProperty('toEnv')) {
                obj['toEnv'] = ApiClient.convertToType(data['toEnv'], 'String');
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
V1IO.prototype['name'] = undefined;

/**
 * @member {String} description
 */
V1IO.prototype['description'] = undefined;

/**
 * @member {String} type
 */
V1IO.prototype['type'] = undefined;

/**
 * @member {Object} value
 */
V1IO.prototype['value'] = undefined;

/**
 * @member {Boolean} isOptional
 */
V1IO.prototype['isOptional'] = undefined;

/**
 * @member {Boolean} isList
 */
V1IO.prototype['isList'] = undefined;

/**
 * @member {Boolean} isFlag
 */
V1IO.prototype['isFlag'] = undefined;

/**
 * @member {String} argFormat
 */
V1IO.prototype['argFormat'] = undefined;

/**
 * @member {Boolean} delayValidation
 */
V1IO.prototype['delayValidation'] = undefined;

/**
 * @member {Array.<Object>} options
 */
V1IO.prototype['options'] = undefined;

/**
 * @member {String} connection
 */
V1IO.prototype['connection'] = undefined;

/**
 * @member {Boolean} toInit
 */
V1IO.prototype['toInit'] = undefined;

/**
 * @member {String} toEnv
 */
V1IO.prototype['toEnv'] = undefined;






export default V1IO;

