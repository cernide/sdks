/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1Mapping model module.
 * @module model/V1Mapping
 * @version 2.0.2
 */
class V1Mapping {
    /**
     * Constructs a new <code>V1Mapping</code>.
     * @alias module:model/V1Mapping
     */
    constructor() {

        V1Mapping.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Mapping</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Mapping} obj Optional instance to populate.
     * @return {module:model/V1Mapping} The populated <code>V1Mapping</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Mapping();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('values')) {
                obj['values'] = ApiClient.convertToType(data['values'], [Object]);
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

    /**
     * Validates the JSON data with respect to <code>V1Mapping</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Mapping</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['kind'] && !(typeof data['kind'] === 'string' || data['kind'] instanceof String)) {
            throw new Error("Expected the field `kind` to be a primitive type in the JSON string but got " + data['kind']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['values'])) {
            throw new Error("Expected the field `values` to be an array in the JSON data but got " + data['values']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['earlyStopping'])) {
            throw new Error("Expected the field `earlyStopping` to be an array in the JSON data but got " + data['earlyStopping']);
        }

        return true;
    }


}



/**
 * @member {String} kind
 * @default 'mapping'
 */
V1Mapping.prototype['kind'] = 'mapping';

/**
 * @member {Array.<Object>} values
 */
V1Mapping.prototype['values'] = undefined;

/**
 * @member {Number} concurrency
 */
V1Mapping.prototype['concurrency'] = undefined;

/**
 * @member {Array.<Object>} earlyStopping
 */
V1Mapping.prototype['earlyStopping'] = undefined;






export default V1Mapping;

