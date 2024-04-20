/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.8
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1PasswordChange model module.
 * @module model/V1PasswordChange
 * @version 2.1.8
 */
class V1PasswordChange {
    /**
     * Constructs a new <code>V1PasswordChange</code>.
     * @alias module:model/V1PasswordChange
     */
    constructor() { 
        
        V1PasswordChange.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1PasswordChange</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1PasswordChange} obj Optional instance to populate.
     * @return {module:model/V1PasswordChange} The populated <code>V1PasswordChange</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1PasswordChange();

            if (data.hasOwnProperty('old_password')) {
                obj['old_password'] = ApiClient.convertToType(data['old_password'], 'String');
            }
            if (data.hasOwnProperty('new_password1')) {
                obj['new_password1'] = ApiClient.convertToType(data['new_password1'], 'String');
            }
            if (data.hasOwnProperty('new_password2')) {
                obj['new_password2'] = ApiClient.convertToType(data['new_password2'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1PasswordChange</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1PasswordChange</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['old_password'] && !(typeof data['old_password'] === 'string' || data['old_password'] instanceof String)) {
            throw new Error("Expected the field `old_password` to be a primitive type in the JSON string but got " + data['old_password']);
        }
        // ensure the json data is a string
        if (data['new_password1'] && !(typeof data['new_password1'] === 'string' || data['new_password1'] instanceof String)) {
            throw new Error("Expected the field `new_password1` to be a primitive type in the JSON string but got " + data['new_password1']);
        }
        // ensure the json data is a string
        if (data['new_password2'] && !(typeof data['new_password2'] === 'string' || data['new_password2'] instanceof String)) {
            throw new Error("Expected the field `new_password2` to be a primitive type in the JSON string but got " + data['new_password2']);
        }

        return true;
    }


}



/**
 * @member {String} old_password
 */
V1PasswordChange.prototype['old_password'] = undefined;

/**
 * @member {String} new_password1
 */
V1PasswordChange.prototype['new_password1'] = undefined;

/**
 * @member {String} new_password2
 */
V1PasswordChange.prototype['new_password2'] = undefined;






export default V1PasswordChange;

