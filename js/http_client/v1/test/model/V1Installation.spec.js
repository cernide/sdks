/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.2-rc1
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
    instance = new PolyaxonSdk.V1Installation();
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

  describe('V1Installation', function() {
    it('should create an instance of V1Installation', function() {
      // uncomment below and update the code to test V1Installation
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be.a(PolyaxonSdk.V1Installation);
    });

    it('should have the property key (base name: "key")', function() {
      // uncomment below and update the code to test the property key
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

    it('should have the property version (base name: "version")', function() {
      // uncomment below and update the code to test the property version
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

    it('should have the property dist (base name: "dist")', function() {
      // uncomment below and update the code to test the property dist
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

    it('should have the property host (base name: "host")', function() {
      // uncomment below and update the code to test the property host
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

    it('should have the property hmac (base name: "hmac")', function() {
      // uncomment below and update the code to test the property hmac
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

    it('should have the property mode (base name: "mode")', function() {
      // uncomment below and update the code to test the property mode
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

    it('should have the property orgs (base name: "orgs")', function() {
      // uncomment below and update the code to test the property orgs
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

    it('should have the property auth (base name: "auth")', function() {
      // uncomment below and update the code to test the property auth
      //var instance = new PolyaxonSdk.V1Installation();
      //expect(instance).to.be();
    });

  });

}));
