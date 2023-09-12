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
    instance = new PolyaxonSdk.SearchesV1Api();
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

  describe('SearchesV1Api', function() {
    describe('createSearch', function() {
      it('should call createSearch successfully', function(done) {
        //uncomment below and update the code to test createSearch
        //instance.createSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteSearch', function() {
      it('should call deleteSearch successfully', function(done) {
        //uncomment below and update the code to test deleteSearch
        //instance.deleteSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getSearch', function() {
      it('should call getSearch successfully', function(done) {
        //uncomment below and update the code to test getSearch
        //instance.getSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listSearchNames', function() {
      it('should call listSearchNames successfully', function(done) {
        //uncomment below and update the code to test listSearchNames
        //instance.listSearchNames(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listSearches', function() {
      it('should call listSearches successfully', function(done) {
        //uncomment below and update the code to test listSearches
        //instance.listSearches(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchSearch', function() {
      it('should call patchSearch successfully', function(done) {
        //uncomment below and update the code to test patchSearch
        //instance.patchSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateSearch', function() {
      it('should call updateSearch successfully', function(done) {
        //uncomment below and update the code to test updateSearch
        //instance.updateSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
