/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
/**
* Enum class V1CloningKind.
* @enum {}
* @readonly
*/
export default class V1CloningKind {

        /**
         * value: "copy"
         * @const
         */
        "copy" = "copy";


        /**
         * value: "restart"
         * @const
         */
        "restart" = "restart";


        /**
         * value: "cache"
         * @const
         */
        "cache" = "cache";



    /**
    * Returns a <code>V1CloningKind</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1CloningKind} The enum <code>V1CloningKind</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

