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
/**
* Enum class V1PatchStrategy.
* @enum {}
* @readonly
*/
export default class V1PatchStrategy {

        /**
         * value: "replace"
         * @const
         */
        "replace" = "replace";


        /**
         * value: "isnull"
         * @const
         */
        "isnull" = "isnull";


        /**
         * value: "post_merge"
         * @const
         */
        "post_merge" = "post_merge";


        /**
         * value: "pre_merge"
         * @const
         */
        "pre_merge" = "pre_merge";



    /**
    * Returns a <code>V1PatchStrategy</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1PatchStrategy} The enum <code>V1PatchStrategy</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

