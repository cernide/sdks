/**
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.0.0-rc16
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
    instance = new PolyaxonSdk.V1MedianStoppingPolicy();
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

  describe('V1MedianStoppingPolicy', function() {
    it('should create an instance of V1MedianStoppingPolicy', function() {
      // uncomment below and update the code to test V1MedianStoppingPolicy
      //var instance = new PolyaxonSdk.V1MedianStoppingPolicy();
      //expect(instance).to.be.a(PolyaxonSdk.V1MedianStoppingPolicy);
    });

    it('should have the property kind (base name: "kind")', function() {
      // uncomment below and update the code to test the property kind
      //var instance = new PolyaxonSdk.V1MedianStoppingPolicy();
      //expect(instance).to.be();
    });

    it('should have the property evaluationInterval (base name: "evaluationInterval")', function() {
      // uncomment below and update the code to test the property evaluationInterval
      //var instance = new PolyaxonSdk.V1MedianStoppingPolicy();
      //expect(instance).to.be();
    });

    it('should have the property minInterval (base name: "minInterval")', function() {
      // uncomment below and update the code to test the property minInterval
      //var instance = new PolyaxonSdk.V1MedianStoppingPolicy();
      //expect(instance).to.be();
    });

    it('should have the property minSamples (base name: "minSamples")', function() {
      // uncomment below and update the code to test the property minSamples
      //var instance = new PolyaxonSdk.V1MedianStoppingPolicy();
      //expect(instance).to.be();
    });

  });

}));
