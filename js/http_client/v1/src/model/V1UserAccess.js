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

/**
 * The V1UserAccess model module.
 * @module model/V1UserAccess
 * @version 2.1.6-rc0
 */
class V1UserAccess {
    /**
     * Constructs a new <code>V1UserAccess</code>.
     * @alias module:model/V1UserAccess
     */
    constructor() { 
        
        V1UserAccess.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1UserAccess</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1UserAccess} obj Optional instance to populate.
     * @return {module:model/V1UserAccess} The populated <code>V1UserAccess</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1UserAccess();

            if (data.hasOwnProperty('user')) {
                obj['user'] = ApiClient.convertToType(data['user'], 'String');
            }
            if (data.hasOwnProperty('queue')) {
                obj['queue'] = ApiClient.convertToType(data['queue'], 'String');
            }
            if (data.hasOwnProperty('preset')) {
                obj['preset'] = ApiClient.convertToType(data['preset'], 'String');
            }
            if (data.hasOwnProperty('namespace')) {
                obj['namespace'] = ApiClient.convertToType(data['namespace'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1UserAccess</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1UserAccess</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['user'] && !(typeof data['user'] === 'string' || data['user'] instanceof String)) {
            throw new Error("Expected the field `user` to be a primitive type in the JSON string but got " + data['user']);
        }
        // ensure the json data is a string
        if (data['queue'] && !(typeof data['queue'] === 'string' || data['queue'] instanceof String)) {
            throw new Error("Expected the field `queue` to be a primitive type in the JSON string but got " + data['queue']);
        }
        // ensure the json data is a string
        if (data['preset'] && !(typeof data['preset'] === 'string' || data['preset'] instanceof String)) {
            throw new Error("Expected the field `preset` to be a primitive type in the JSON string but got " + data['preset']);
        }
        // ensure the json data is a string
        if (data['namespace'] && !(typeof data['namespace'] === 'string' || data['namespace'] instanceof String)) {
            throw new Error("Expected the field `namespace` to be a primitive type in the JSON string but got " + data['namespace']);
        }

        return true;
    }


}



/**
 * @member {String} user
 */
V1UserAccess.prototype['user'] = undefined;

/**
 * @member {String} queue
 */
V1UserAccess.prototype['queue'] = undefined;

/**
 * @member {String} preset
 */
V1UserAccess.prototype['preset'] = undefined;

/**
 * @member {String} namespace
 */
V1UserAccess.prototype['namespace'] = undefined;






export default V1UserAccess;

