/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.6-rc0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1Component from './V1Component';
import V1Environment from './V1Environment';
import V1Operation from './V1Operation';

/**
 * The V1Dag model module.
 * @module model/V1Dag
 * @version 2.1.6-rc0
 */
class V1Dag {
    /**
     * Constructs a new <code>V1Dag</code>.
     * @alias module:model/V1Dag
     */
    constructor() { 
        
        V1Dag.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Dag</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Dag} obj Optional instance to populate.
     * @return {module:model/V1Dag} The populated <code>V1Dag</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Dag();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('operations')) {
                obj['operations'] = ApiClient.convertToType(data['operations'], [V1Operation]);
            }
            if (data.hasOwnProperty('components')) {
                obj['components'] = ApiClient.convertToType(data['components'], [V1Component]);
            }
            if (data.hasOwnProperty('concurrency')) {
                obj['concurrency'] = ApiClient.convertToType(data['concurrency'], 'Number');
            }
            if (data.hasOwnProperty('earlyStopping')) {
                obj['earlyStopping'] = ApiClient.convertToType(data['earlyStopping'], [Object]);
            }
            if (data.hasOwnProperty('environment')) {
                obj['environment'] = V1Environment.constructFromObject(data['environment']);
            }
            if (data.hasOwnProperty('connections')) {
                obj['connections'] = ApiClient.convertToType(data['connections'], ['String']);
            }
            if (data.hasOwnProperty('volumes')) {
                obj['volumes'] = ApiClient.convertToType(data['volumes'], [Object]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Dag</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Dag</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['kind'] && !(typeof data['kind'] === 'string' || data['kind'] instanceof String)) {
            throw new Error("Expected the field `kind` to be a primitive type in the JSON string but got " + data['kind']);
        }
        if (data['operations']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['operations'])) {
                throw new Error("Expected the field `operations` to be an array in the JSON data but got " + data['operations']);
            }
            // validate the optional field `operations` (array)
            for (const item of data['operations']) {
                V1Operation.validateJSON(item);
            };
        }
        if (data['components']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['components'])) {
                throw new Error("Expected the field `components` to be an array in the JSON data but got " + data['components']);
            }
            // validate the optional field `components` (array)
            for (const item of data['components']) {
                V1Component.validateJSON(item);
            };
        }
        // ensure the json data is an array
        if (!Array.isArray(data['earlyStopping'])) {
            throw new Error("Expected the field `earlyStopping` to be an array in the JSON data but got " + data['earlyStopping']);
        }
        // validate the optional field `environment`
        if (data['environment']) { // data not null
          V1Environment.validateJSON(data['environment']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['connections'])) {
            throw new Error("Expected the field `connections` to be an array in the JSON data but got " + data['connections']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['volumes'])) {
            throw new Error("Expected the field `volumes` to be an array in the JSON data but got " + data['volumes']);
        }

        return true;
    }


}



/**
 * @member {String} kind
 * @default 'dag'
 */
V1Dag.prototype['kind'] = 'dag';

/**
 * @member {Array.<module:model/V1Operation>} operations
 */
V1Dag.prototype['operations'] = undefined;

/**
 * @member {Array.<module:model/V1Component>} components
 */
V1Dag.prototype['components'] = undefined;

/**
 * @member {Number} concurrency
 */
V1Dag.prototype['concurrency'] = undefined;

/**
 * @member {Array.<Object>} earlyStopping
 */
V1Dag.prototype['earlyStopping'] = undefined;

/**
 * @member {module:model/V1Environment} environment
 */
V1Dag.prototype['environment'] = undefined;

/**
 * @member {Array.<String>} connections
 */
V1Dag.prototype['connections'] = undefined;

/**
 * Volumes is a list of volumes that can be mounted.
 * @member {Array.<Object>} volumes
 */
V1Dag.prototype['volumes'] = undefined;






export default V1Dag;

