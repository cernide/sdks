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
import V1Cache from './V1Cache';
import V1Param from './V1Param';
import V1PatchStrategy from './V1PatchStrategy';

/**
 * The V1Build model module.
 * @module model/V1Build
 * @version 2.1.5
 */
class V1Build {
    /**
     * Constructs a new <code>V1Build</code>.
     * @alias module:model/V1Build
     */
    constructor() { 
        
        V1Build.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Build</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Build} obj Optional instance to populate.
     * @return {module:model/V1Build} The populated <code>V1Build</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Build();

            if (data.hasOwnProperty('hubRef')) {
                obj['hubRef'] = ApiClient.convertToType(data['hubRef'], 'String');
            }
            if (data.hasOwnProperty('connection')) {
                obj['connection'] = ApiClient.convertToType(data['connection'], 'String');
            }
            if (data.hasOwnProperty('presets')) {
                obj['presets'] = ApiClient.convertToType(data['presets'], ['String']);
            }
            if (data.hasOwnProperty('queue')) {
                obj['queue'] = ApiClient.convertToType(data['queue'], 'String');
            }
            if (data.hasOwnProperty('namespace')) {
                obj['namespace'] = ApiClient.convertToType(data['namespace'], 'String');
            }
            if (data.hasOwnProperty('cache')) {
                obj['cache'] = V1Cache.constructFromObject(data['cache']);
            }
            if (data.hasOwnProperty('params')) {
                obj['params'] = ApiClient.convertToType(data['params'], {'String': V1Param});
            }
            if (data.hasOwnProperty('runPatch')) {
                obj['runPatch'] = ApiClient.convertToType(data['runPatch'], Object);
            }
            if (data.hasOwnProperty('patchStrategy')) {
                obj['patchStrategy'] = V1PatchStrategy.constructFromObject(data['patchStrategy']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Build</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Build</code>.
     */
    static validateJSON(data) {
        // ensure the json data is a string
        if (data['hubRef'] && !(typeof data['hubRef'] === 'string' || data['hubRef'] instanceof String)) {
            throw new Error("Expected the field `hubRef` to be a primitive type in the JSON string but got " + data['hubRef']);
        }
        // ensure the json data is a string
        if (data['connection'] && !(typeof data['connection'] === 'string' || data['connection'] instanceof String)) {
            throw new Error("Expected the field `connection` to be a primitive type in the JSON string but got " + data['connection']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['presets'])) {
            throw new Error("Expected the field `presets` to be an array in the JSON data but got " + data['presets']);
        }
        // ensure the json data is a string
        if (data['queue'] && !(typeof data['queue'] === 'string' || data['queue'] instanceof String)) {
            throw new Error("Expected the field `queue` to be a primitive type in the JSON string but got " + data['queue']);
        }
        // ensure the json data is a string
        if (data['namespace'] && !(typeof data['namespace'] === 'string' || data['namespace'] instanceof String)) {
            throw new Error("Expected the field `namespace` to be a primitive type in the JSON string but got " + data['namespace']);
        }
        // validate the optional field `cache`
        if (data['cache']) { // data not null
          V1Cache.validateJSON(data['cache']);
        }

        return true;
    }


}



/**
 * @member {String} hubRef
 */
V1Build.prototype['hubRef'] = undefined;

/**
 * @member {String} connection
 */
V1Build.prototype['connection'] = undefined;

/**
 * @member {Array.<String>} presets
 */
V1Build.prototype['presets'] = undefined;

/**
 * @member {String} queue
 */
V1Build.prototype['queue'] = undefined;

/**
 * @member {String} namespace
 */
V1Build.prototype['namespace'] = undefined;

/**
 * @member {module:model/V1Cache} cache
 */
V1Build.prototype['cache'] = undefined;

/**
 * @member {Object.<String, module:model/V1Param>} params
 */
V1Build.prototype['params'] = undefined;

/**
 * @member {Object} runPatch
 */
V1Build.prototype['runPatch'] = undefined;

/**
 * @member {module:model/V1PatchStrategy} patchStrategy
 */
V1Build.prototype['patchStrategy'] = undefined;






export default V1Build;

