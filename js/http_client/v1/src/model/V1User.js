/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.2.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1User model module.
 * @module model/V1User
 * @version 2.2.0
 */
class V1User {
    /**
     * Constructs a new <code>V1User</code>.
     * @alias module:model/V1User
     */
    constructor() {

        V1User.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1User</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1User} obj Optional instance to populate.
     * @return {module:model/V1User} The populated <code>V1User</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1User();

            if (data.hasOwnProperty('username')) {
                obj['username'] = ApiClient.convertToType(data['username'], 'String');
            }
            if (data.hasOwnProperty('email')) {
                obj['email'] = ApiClient.convertToType(data['email'], 'String');
            }
            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('theme')) {
                obj['theme'] = ApiClient.convertToType(data['theme'], 'Number');
            }
            if (data.hasOwnProperty('organization')) {
                obj['organization'] = ApiClient.convertToType(data['organization'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1User</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1User</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['username'] && !(typeof data['username'] === 'string' || data['username'] instanceof String)) {
            throw new Error("Expected the field `username` to be a primitive type in the JSON string but got " + data['username']);
        }
        // ensure the json data is a string
        if (data['email'] && !(typeof data['email'] === 'string' || data['email'] instanceof String)) {
            throw new Error("Expected the field `email` to be a primitive type in the JSON string but got " + data['email']);
        }
        // ensure the json data is a string
        if (data['name'] && !(typeof data['name'] === 'string' || data['name'] instanceof String)) {
            throw new Error("Expected the field `name` to be a primitive type in the JSON string but got " + data['name']);
        }
        // ensure the json data is a string
        if (data['kind'] && !(typeof data['kind'] === 'string' || data['kind'] instanceof String)) {
            throw new Error("Expected the field `kind` to be a primitive type in the JSON string but got " + data['kind']);
        }
        // ensure the json data is a string
        if (data['organization'] && !(typeof data['organization'] === 'string' || data['organization'] instanceof String)) {
            throw new Error("Expected the field `organization` to be a primitive type in the JSON string but got " + data['organization']);
        }

        return true;
    }


}



/**
 * @member {String} username
 */
V1User.prototype['username'] = undefined;

/**
 * @member {String} email
 */
V1User.prototype['email'] = undefined;

/**
 * @member {String} name
 */
V1User.prototype['name'] = undefined;

/**
 * @member {String} kind
 */
V1User.prototype['kind'] = undefined;

/**
 * @member {Number} theme
 */
V1User.prototype['theme'] = undefined;

/**
 * @member {String} organization
 */
V1User.prototype['organization'] = undefined;






export default V1User;

