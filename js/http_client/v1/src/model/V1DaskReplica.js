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
import V1Environment from './V1Environment';
import V1Init from './V1Init';

/**
 * The V1DaskReplica model module.
 * @module model/V1DaskReplica
 * @version 2.1.8
 */
class V1DaskReplica {
    /**
     * Constructs a new <code>V1DaskReplica</code>.
     * @alias module:model/V1DaskReplica
     */
    constructor() { 
        
        V1DaskReplica.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1DaskReplica</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1DaskReplica} obj Optional instance to populate.
     * @return {module:model/V1DaskReplica} The populated <code>V1DaskReplica</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1DaskReplica();

            if (data.hasOwnProperty('replicas')) {
                obj['replicas'] = ApiClient.convertToType(data['replicas'], 'Number');
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
            if (data.hasOwnProperty('init')) {
                obj['init'] = ApiClient.convertToType(data['init'], [V1Init]);
            }
            if (data.hasOwnProperty('sidecars')) {
                obj['sidecars'] = ApiClient.convertToType(data['sidecars'], [Object]);
            }
            if (data.hasOwnProperty('container')) {
                obj['container'] = ApiClient.convertToType(data['container'], Object);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1DaskReplica</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1DaskReplica</code>.
     */
    static validateJSON(data) {
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
        if (data['init']) { // data not null
            // ensure the json data is an array
            if (!Array.isArray(data['init'])) {
                throw new Error("Expected the field `init` to be an array in the JSON data but got " + data['init']);
            }
            // validate the optional field `init` (array)
            for (const item of data['init']) {
                V1Init.validateJSON(item);
            };
        }
        // ensure the json data is an array
        if (!Array.isArray(data['sidecars'])) {
            throw new Error("Expected the field `sidecars` to be an array in the JSON data but got " + data['sidecars']);
        }

        return true;
    }


}



/**
 * @member {Number} replicas
 */
V1DaskReplica.prototype['replicas'] = undefined;

/**
 * @member {module:model/V1Environment} environment
 */
V1DaskReplica.prototype['environment'] = undefined;

/**
 * @member {Array.<String>} connections
 */
V1DaskReplica.prototype['connections'] = undefined;

/**
 * @member {Array.<Object>} volumes
 */
V1DaskReplica.prototype['volumes'] = undefined;

/**
 * @member {Array.<module:model/V1Init>} init
 */
V1DaskReplica.prototype['init'] = undefined;

/**
 * @member {Array.<Object>} sidecars
 */
V1DaskReplica.prototype['sidecars'] = undefined;

/**
 * @member {Object} container
 */
V1DaskReplica.prototype['container'] = undefined;






export default V1DaskReplica;

