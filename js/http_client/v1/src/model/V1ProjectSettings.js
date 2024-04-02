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
import V1UserAccess from './V1UserAccess';

/**
 * The V1ProjectSettings model module.
 * @module model/V1ProjectSettings
 * @version 2.1.5
 */
class V1ProjectSettings {
    /**
     * Constructs a new <code>V1ProjectSettings</code>.
     * @alias module:model/V1ProjectSettings
     */
    constructor() { 
        
        V1ProjectSettings.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1ProjectSettings</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1ProjectSettings} obj Optional instance to populate.
     * @return {module:model/V1ProjectSettings} The populated <code>V1ProjectSettings</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1ProjectSettings();

            if (data.hasOwnProperty('connections')) {
                obj['connections'] = ApiClient.convertToType(data['connections'], ['String']);
            }
            if (data.hasOwnProperty('preset')) {
                obj['preset'] = ApiClient.convertToType(data['preset'], 'String');
            }
            if (data.hasOwnProperty('presets')) {
                obj['presets'] = ApiClient.convertToType(data['presets'], ['String']);
            }
            if (data.hasOwnProperty('queue')) {
                obj['queue'] = ApiClient.convertToType(data['queue'], 'String');
            }
            if (data.hasOwnProperty('queues')) {
                obj['queues'] = ApiClient.convertToType(data['queues'], ['String']);
            }
            if (data.hasOwnProperty('agents')) {
                obj['agents'] = ApiClient.convertToType(data['agents'], ['String']);
            }
            if (data.hasOwnProperty('namespaces')) {
                obj['namespaces'] = ApiClient.convertToType(data['namespaces'], ['String']);
            }
            if (data.hasOwnProperty('user_accesses')) {
                obj['user_accesses'] = ApiClient.convertToType(data['user_accesses'], [V1UserAccess]);
            }
            if (data.hasOwnProperty('teams')) {
                obj['teams'] = ApiClient.convertToType(data['teams'], ['String']);
            }
            if (data.hasOwnProperty('projects')) {
                obj['projects'] = ApiClient.convertToType(data['projects'], ['String']);
            }
            if (data.hasOwnProperty('policy')) {
                obj['policy'] = ApiClient.convertToType(data['policy'], 'String');
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1ProjectSettings</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1ProjectSettings</code>.
     */
    static validateJSON(data) {
        // ensure the json data is an array
        if (!Array.isArray(data['connections'])) {
            throw new Error("Expected the field `connections` to be an array in the JSON data but got " + data['connections']);
        }
        // ensure the json data is a string
        if (data['preset'] && !(typeof data['preset'] === 'string' || data['preset'] instanceof String)) {
            throw new Error("Expected the field `preset` to be a primitive type in the JSON string but got " + data['preset']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['presets'])) {
            throw new Error("Expected the field `presets` to be an array in the JSON data but got " + data['presets']);
        }
        // ensure the json data is a string
        if (data['queue'] && !(typeof data['queue'] === 'string' || data['queue'] instanceof String)) {
            throw new Error("Expected the field `queue` to be a primitive type in the JSON string but got " + data['queue']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['queues'])) {
            throw new Error("Expected the field `queues` to be an array in the JSON data but got " + data['queues']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['agents'])) {
            throw new Error("Expected the field `agents` to be an array in the JSON data but got " + data['agents']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['namespaces'])) {
            throw new Error("Expected the field `namespaces` to be an array in the JSON data but got " + data['namespaces']);
        }
        if (data['user_accesses']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['user_accesses'])) {
                throw new Error("Expected the field `user_accesses` to be an array in the JSON data but got " + data['user_accesses']);
            }
            // validate the optional field `user_accesses` (array)
            for (const item of data['user_accesses']) {
                V1UserAccess.validateJSON(item);
            };
        }
        // ensure the json data is an array
        if (!Array.isArray(data['teams'])) {
            throw new Error("Expected the field `teams` to be an array in the JSON data but got " + data['teams']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['projects'])) {
            throw new Error("Expected the field `projects` to be an array in the JSON data but got " + data['projects']);
        }
        // ensure the json data is a string
        if (data['policy'] && !(typeof data['policy'] === 'string' || data['policy'] instanceof String)) {
            throw new Error("Expected the field `policy` to be a primitive type in the JSON string but got " + data['policy']);
        }

        return true;
    }


}



/**
 * @member {Array.<String>} connections
 */
V1ProjectSettings.prototype['connections'] = undefined;

/**
 * @member {String} preset
 */
V1ProjectSettings.prototype['preset'] = undefined;

/**
 * @member {Array.<String>} presets
 */
V1ProjectSettings.prototype['presets'] = undefined;

/**
 * @member {String} queue
 */
V1ProjectSettings.prototype['queue'] = undefined;

/**
 * @member {Array.<String>} queues
 */
V1ProjectSettings.prototype['queues'] = undefined;

/**
 * @member {Array.<String>} agents
 */
V1ProjectSettings.prototype['agents'] = undefined;

/**
 * @member {Array.<String>} namespaces
 */
V1ProjectSettings.prototype['namespaces'] = undefined;

/**
 * @member {Array.<module:model/V1UserAccess>} user_accesses
 */
V1ProjectSettings.prototype['user_accesses'] = undefined;

/**
 * @member {Array.<String>} teams
 */
V1ProjectSettings.prototype['teams'] = undefined;

/**
 * @member {Array.<String>} projects
 */
V1ProjectSettings.prototype['projects'] = undefined;

/**
 * @member {String} policy
 */
V1ProjectSettings.prototype['policy'] = undefined;






export default V1ProjectSettings;

