/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
/**
* Enum class V1ResourceType.
* @enum {}
* @readonly
*/
export default class V1ResourceType {

        /**
         * value: "int"
         * @const
         */
        "int" = "int";


        /**
         * value: "float"
         * @const
         */
        "float" = "float";



    /**
    * Returns a <code>V1ResourceType</code> enum value from a Javascript object name.
    * @param {Object} data The plain JavaScript object containing the name of the enum value.
    * @return {module:model/V1ResourceType} The enum <code>V1ResourceType</code> value.
    */
    static constructFromObject(object) {
        return object;
    }
}

