/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc30
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';

/**
 * The V1Activity model module.
 * @module model/V1Activity
 * @version 2.0.0-rc30
 */
class V1Activity {
    /**
     * Constructs a new <code>V1Activity</code>.
     * @alias module:model/V1Activity
     */
    constructor() {

        V1Activity.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Activity</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Activity} obj Optional instance to populate.
     * @return {module:model/V1Activity} The populated <code>V1Activity</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Activity();

            if (data.hasOwnProperty('actor')) {
                obj['actor'] = ApiClient.convertToType(data['actor'], 'String');
            }
            if (data.hasOwnProperty('owner')) {
                obj['owner'] = ApiClient.convertToType(data['owner'], 'String');
            }
            if (data.hasOwnProperty('created_at')) {
                obj['created_at'] = ApiClient.convertToType(data['created_at'], 'Date');
            }
            if (data.hasOwnProperty('event_action')) {
                obj['event_action'] = ApiClient.convertToType(data['event_action'], 'String');
            }
            if (data.hasOwnProperty('event_subject')) {
                obj['event_subject'] = ApiClient.convertToType(data['event_subject'], 'String');
            }
            if (data.hasOwnProperty('object_name')) {
                obj['object_name'] = ApiClient.convertToType(data['object_name'], 'String');
            }
            if (data.hasOwnProperty('object_uuid')) {
                obj['object_uuid'] = ApiClient.convertToType(data['object_uuid'], 'String');
            }
            if (data.hasOwnProperty('object_parent')) {
                obj['object_parent'] = ApiClient.convertToType(data['object_parent'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Activity</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Activity</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['actor'] && !(typeof data['actor'] === 'string' || data['actor'] instanceof String)) {
            throw new Error("Expected the field `actor` to be a primitive type in the JSON string but got " + data['actor']);
        }
        // ensure the json data is a string
        if (data['owner'] && !(typeof data['owner'] === 'string' || data['owner'] instanceof String)) {
            throw new Error("Expected the field `owner` to be a primitive type in the JSON string but got " + data['owner']);
        }
        // ensure the json data is a string
        if (data['event_action'] && !(typeof data['event_action'] === 'string' || data['event_action'] instanceof String)) {
            throw new Error("Expected the field `event_action` to be a primitive type in the JSON string but got " + data['event_action']);
        }
        // ensure the json data is a string
        if (data['event_subject'] && !(typeof data['event_subject'] === 'string' || data['event_subject'] instanceof String)) {
            throw new Error("Expected the field `event_subject` to be a primitive type in the JSON string but got " + data['event_subject']);
        }
        // ensure the json data is a string
        if (data['object_name'] && !(typeof data['object_name'] === 'string' || data['object_name'] instanceof String)) {
            throw new Error("Expected the field `object_name` to be a primitive type in the JSON string but got " + data['object_name']);
        }
        // ensure the json data is a string
        if (data['object_uuid'] && !(typeof data['object_uuid'] === 'string' || data['object_uuid'] instanceof String)) {
            throw new Error("Expected the field `object_uuid` to be a primitive type in the JSON string but got " + data['object_uuid']);
        }
        // ensure the json data is a string
        if (data['object_parent'] && !(typeof data['object_parent'] === 'string' || data['object_parent'] instanceof String)) {
            throw new Error("Expected the field `object_parent` to be a primitive type in the JSON string but got " + data['object_parent']);
        }

        return true;
    }


}



/**
 * @member {String} actor
 */
V1Activity.prototype['actor'] = undefined;

/**
 * @member {String} owner
 */
V1Activity.prototype['owner'] = undefined;

/**
 * @member {Date} created_at
 */
V1Activity.prototype['created_at'] = undefined;

/**
 * @member {String} event_action
 */
V1Activity.prototype['event_action'] = undefined;

/**
 * @member {String} event_subject
 */
V1Activity.prototype['event_subject'] = undefined;

/**
 * @member {String} object_name
 */
V1Activity.prototype['object_name'] = undefined;

/**
 * @member {String} object_uuid
 */
V1Activity.prototype['object_uuid'] = undefined;

/**
 * @member {String} object_parent
 */
V1Activity.prototype['object_parent'] = undefined;






export default V1Activity;

