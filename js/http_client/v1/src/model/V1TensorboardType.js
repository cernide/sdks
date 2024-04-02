/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1TensorboardType model module.
 * @module model/V1TensorboardType
 * @version 2.1.5
 */
class V1TensorboardType {
    /**
     * Constructs a new <code>V1TensorboardType</code>.
     * @alias module:model/V1TensorboardType
     */
    constructor() { 
        
        V1TensorboardType.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1TensorboardType</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1TensorboardType} obj Optional instance to populate.
     * @return {module:model/V1TensorboardType} The populated <code>V1TensorboardType</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1TensorboardType();

            if (data.hasOwnProperty('port')) {
                obj['port'] = ApiClient.convertToType(data['port'], 'Number');
            }
            if (data.hasOwnProperty('uuids')) {
                obj['uuids'] = ApiClient.convertToType(data['uuids'], ['String']);
            }
            if (data.hasOwnProperty('use_names')) {
                obj['use_names'] = ApiClient.convertToType(data['use_names'], 'Boolean');
            }
            if (data.hasOwnProperty('path_prefix')) {
                obj['path_prefix'] = ApiClient.convertToType(data['path_prefix'], 'String');
            }
            if (data.hasOwnProperty('plugins')) {
                obj['plugins'] = ApiClient.convertToType(data['plugins'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1TensorboardType</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1TensorboardType</code>.
     */
    static validateJSON(data) {
        // ensure the json data is an array
        if (!Array.isArray(data['uuids'])) {
            throw new Error("Expected the field `uuids` to be an array in the JSON data but got " + data['uuids']);
        }
        // ensure the json data is a string
        if (data['path_prefix'] && !(typeof data['path_prefix'] === 'string' || data['path_prefix'] instanceof String)) {
            throw new Error("Expected the field `path_prefix` to be a primitive type in the JSON string but got " + data['path_prefix']);
        }
        // ensure the json data is a string
        if (data['plugins'] && !(typeof data['plugins'] === 'string' || data['plugins'] instanceof String)) {
            throw new Error("Expected the field `plugins` to be a primitive type in the JSON string but got " + data['plugins']);
        }

        return true;
    }


}



/**
 * @member {Number} port
 */
V1TensorboardType.prototype['port'] = undefined;

/**
 * @member {Array.<String>} uuids
 */
V1TensorboardType.prototype['uuids'] = undefined;

/**
 * @member {Boolean} use_names
 */
V1TensorboardType.prototype['use_names'] = undefined;

/**
 * @member {String} path_prefix
 */
V1TensorboardType.prototype['path_prefix'] = undefined;

/**
 * @member {String} plugins
 */
V1TensorboardType.prototype['plugins'] = undefined;






export default V1TensorboardType;

