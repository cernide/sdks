/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.4
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
    instance = new PolyaxonSdk.V1HpParams();
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

  describe('V1HpParams', function() {
    it('should create an instance of V1HpParams', function() {
      // uncomment below and update the code to test V1HpParams
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be.a(PolyaxonSdk.V1HpParams);
    });

    it('should have the property choice (base name: "choice")', function() {
      // uncomment below and update the code to test the property choice
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property pchoice (base name: "pchoice")', function() {
      // uncomment below and update the code to test the property pchoice
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property range (base name: "range")', function() {
      // uncomment below and update the code to test the property range
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property linspace (base name: "linspace")', function() {
      // uncomment below and update the code to test the property linspace
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property logspace (base name: "logspace")', function() {
      // uncomment below and update the code to test the property logspace
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property geomspace (base name: "geomspace")', function() {
      // uncomment below and update the code to test the property geomspace
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property uniform (base name: "uniform")', function() {
      // uncomment below and update the code to test the property uniform
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property quniform (base name: "quniform")', function() {
      // uncomment below and update the code to test the property quniform
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property loguniform (base name: "loguniform")', function() {
      // uncomment below and update the code to test the property loguniform
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property qloguniform (base name: "qloguniform")', function() {
      // uncomment below and update the code to test the property qloguniform
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property normal (base name: "normal")', function() {
      // uncomment below and update the code to test the property normal
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property qnormal (base name: "qnormal")', function() {
      // uncomment below and update the code to test the property qnormal
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property lognormal (base name: "lognormal")', function() {
      // uncomment below and update the code to test the property lognormal
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property qlognormal (base name: "qlognormal")', function() {
      // uncomment below and update the code to test the property qlognormal
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property daterange (base name: "daterange")', function() {
      // uncomment below and update the code to test the property daterange
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

    it('should have the property datetimerange (base name: "datetimerange")', function() {
      // uncomment below and update the code to test the property datetimerange
      //var instance = new PolyaxonSdk.V1HpParams();
      //expect(instance).to.be();
    });

  });

}));
