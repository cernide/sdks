/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.4.2
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function (root, factory) {
    if (typeof define === "function" && define.amd) {
        // AMD.
        define(["expect.js", process.cwd() + "/src/index"], factory);
    } else if (typeof module === "object" && module.exports) {
        // CommonJS-like environments that support module.exports, like Node.
        factory(require("expect.js"), require(process.cwd() + "/src/index"));
    } else {
        // Browser globals (root is window)
        factory(root.expect, root.PolyaxonSdk);
    }
})(this, function (expect, PolyaxonSdk) {
    "use strict";

    var instance;

    beforeEach(function () {
        instance = new PolyaxonSdk.V1EntitiesTags();
    });

    var getProperty = function (object, getter, property) {
        // Use getter method if present; otherwise, get the property directly.
        if (typeof object[getter] === "function") return object[getter]();
        else return object[property];
    };

    var setProperty = function (object, setter, property, value) {
        // Use setter method if present; otherwise, set the property directly.
        if (typeof object[setter] === "function") object[setter](value);
        else object[property] = value;
    };

    describe("V1EntitiesTags", function () {
        it("should create an instance of V1EntitiesTags", function () {
            // uncomment below and update the code to test V1EntitiesTags
            //var instance = new PolyaxonSdk.V1EntitiesTags();
            //expect(instance).to.be.a(PolyaxonSdk.V1EntitiesTags);
        });

        it('should have the property uuids (base name: "uuids")', function () {
            // uncomment below and update the code to test the property uuids
            //var instance = new PolyaxonSdk.V1EntitiesTags();
            //expect(instance).to.be();
        });

        it('should have the property tags (base name: "tags")', function () {
            // uncomment below and update the code to test the property tags
            //var instance = new PolyaxonSdk.V1EntitiesTags();
            //expect(instance).to.be();
        });
    });
});
