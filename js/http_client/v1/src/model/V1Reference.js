/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc41
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1DagRef from './V1DagRef';
import V1HubRef from './V1HubRef';
import V1PathRef from './V1PathRef';
import V1UrlRef from './V1UrlRef';

/**
 * The V1Reference model module.
 * @module model/V1Reference
 * @version 2.0.0-rc41
 */
class V1Reference {
    /**
     * Constructs a new <code>V1Reference</code>.
     * @alias module:model/V1Reference
     */
    constructor() {

        V1Reference.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) {
    }

    /**
     * Constructs a <code>V1Reference</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Reference} obj Optional instance to populate.
     * @return {module:model/V1Reference} The populated <code>V1Reference</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Reference();

            if (data.hasOwnProperty('hubRef')) {
                obj['hubRef'] = V1HubRef.constructFromObject(data['hubRef']);
            }
            if (data.hasOwnProperty('dagRef')) {
                obj['dagRef'] = V1DagRef.constructFromObject(data['dagRef']);
            }
            if (data.hasOwnProperty('urlRef')) {
                obj['urlRef'] = V1UrlRef.constructFromObject(data['urlRef']);
            }
            if (data.hasOwnProperty('pathRef')) {
                obj['pathRef'] = V1PathRef.constructFromObject(data['pathRef']);
            }
        }
        return obj;
    }

    /**
     * Validates the JSON data with respect to <code>V1Reference</code>.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @return {boolean} to indicate whether the JSON data is valid with respect to <code>V1Reference</code>.
     */
    static validateJSON(data) {
        // validate the optional field `hubRef`
        if (data['hubRef']) { // data not null
          V1HubRef.validateJSON(data['hubRef']);
        }
        // validate the optional field `dagRef`
        if (data['dagRef']) { // data not null
          V1DagRef.validateJSON(data['dagRef']);
        }
        // validate the optional field `urlRef`
        if (data['urlRef']) { // data not null
          V1UrlRef.validateJSON(data['urlRef']);
        }
        // validate the optional field `pathRef`
        if (data['pathRef']) { // data not null
          V1PathRef.validateJSON(data['pathRef']);
        }

        return true;
    }


}



/**
 * @member {module:model/V1HubRef} hubRef
 */
V1Reference.prototype['hubRef'] = undefined;

/**
 * @member {module:model/V1DagRef} dagRef
 */
V1Reference.prototype['dagRef'] = undefined;

/**
 * @member {module:model/V1UrlRef} urlRef
 */
V1Reference.prototype['urlRef'] = undefined;

/**
 * @member {module:model/V1PathRef} pathRef
 */
V1Reference.prototype['pathRef'] = undefined;






export default V1Reference;

