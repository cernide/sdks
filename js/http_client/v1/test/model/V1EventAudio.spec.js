/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc21
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
    instance = new PolyaxonSdk.V1EventAudio();
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

  describe('V1EventAudio', function() {
    it('should create an instance of V1EventAudio', function() {
      // uncomment below and update the code to test V1EventAudio
      //var instance = new PolyaxonSdk.V1EventAudio();
      //expect(instance).to.be.a(PolyaxonSdk.V1EventAudio);
    });

    it('should have the property sample_rate (base name: "sample_rate")', function() {
      // uncomment below and update the code to test the property sample_rate
      //var instance = new PolyaxonSdk.V1EventAudio();
      //expect(instance).to.be();
    });

    it('should have the property num_channels (base name: "num_channels")', function() {
      // uncomment below and update the code to test the property num_channels
      //var instance = new PolyaxonSdk.V1EventAudio();
      //expect(instance).to.be();
    });

    it('should have the property length_frames (base name: "length_frames")', function() {
      // uncomment below and update the code to test the property length_frames
      //var instance = new PolyaxonSdk.V1EventAudio();
      //expect(instance).to.be();
    });

    it('should have the property content_type (base name: "content_type")', function() {
      // uncomment below and update the code to test the property content_type
      //var instance = new PolyaxonSdk.V1EventAudio();
      //expect(instance).to.be();
    });

    it('should have the property path (base name: "path")', function() {
      // uncomment below and update the code to test the property path
      //var instance = new PolyaxonSdk.V1EventAudio();
      //expect(instance).to.be();
    });

  });

}));
