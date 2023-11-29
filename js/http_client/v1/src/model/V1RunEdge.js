/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1Run from './V1Run';
import V1RunEdgeKind from './V1RunEdgeKind';
import V1Statuses from './V1Statuses';

/**
 * The V1RunEdge model module.
 * @module model/V1RunEdge
 * @version 2.0.3
 */
class V1RunEdge {
    /**
     * Constructs a new <code>V1RunEdge</code>.
     * @alias module:model/V1RunEdge
     */
    constructor() { 
        
        V1RunEdge.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1RunEdge</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1RunEdge} obj Optional instance to populate.
     * @return {module:model/V1RunEdge} The populated <code>V1RunEdge</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1RunEdge();

            if (data.hasOwnProperty('upstream')) {
                obj['upstream'] = V1Run.constructFromObject(data['upstream']);
            }
            if (data.hasOwnProperty('downstream')) {
                obj['downstream'] = V1Run.constructFromObject(data['downstream']);
            }
            if (data.hasOwnProperty('kind')) {
                obj['kind'] = V1RunEdgeKind.constructFromObject(data['kind']);
            }
            if (data.hasOwnProperty('values')) {
                obj['values'] = ApiClient.convertToType(data['values'], Object);
            }
            if (data.hasOwnProperty('statuses')) {
                obj['statuses'] = ApiClient.convertToType(data['statuses'], [V1Statuses]);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1RunEdge</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1RunEdge</code>.
     */
    static validateJSON(data) {
        // validate the optional field `upstream`
        if (data['upstream']) { // data not null
          V1Run.validateJSON(data['upstream']);
        }
        // validate the optional field `downstream`
        if (data['downstream']) { // data not null
          V1Run.validateJSON(data['downstream']);
        }
        // ensure the json data is an array
        if (!Array.isArray(data['statuses'])) {
            throw new Error("Expected the field `statuses` to be an array in the JSON data but got " + data['statuses']);
        }

        return true;
    }


}



/**
 * @member {module:model/V1Run} upstream
 */
V1RunEdge.prototype['upstream'] = undefined;

/**
 * @member {module:model/V1Run} downstream
 */
V1RunEdge.prototype['downstream'] = undefined;

/**
 * @member {module:model/V1RunEdgeKind} kind
 */
V1RunEdge.prototype['kind'] = undefined;

/**
 * @member {Object} values
 */
V1RunEdge.prototype['values'] = undefined;

/**
 * @member {Array.<module:model/V1Statuses>} statuses
 */
V1RunEdge.prototype['statuses'] = undefined;






export default V1RunEdge;

