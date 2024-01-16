/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.0-rc5
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD.
    define(['expect.js', process.cwd()+'/src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require(process.cwd()+'/src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.PolyaxonSdk);
  }
}(this, function(expect, PolyaxonSdk) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new PolyaxonSdk.V1EventConfusionMatrix();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('V1EventConfusionMatrix', function() {
    it('should create an instance of V1EventConfusionMatrix', function() {
      // uncomment below and update the code to test V1EventConfusionMatrix
      //var instance = new PolyaxonSdk.V1EventConfusionMatrix();
      //expect(instance).to.be.a(PolyaxonSdk.V1EventConfusionMatrix);
    });

    it('should have the property x (base name: "x")', function() {
      // uncomment below and update the code to test the property x
      //var instance = new PolyaxonSdk.V1EventConfusionMatrix();
      //expect(instance).to.be();
    });

    it('should have the property y (base name: "y")', function() {
      // uncomment below and update the code to test the property y
      //var instance = new PolyaxonSdk.V1EventConfusionMatrix();
      //expect(instance).to.be();
    });

    it('should have the property z (base name: "z")', function() {
      // uncomment below and update the code to test the property z
      //var instance = new PolyaxonSdk.V1EventConfusionMatrix();
      //expect(instance).to.be();
    });

  });

}));
