/**
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc30
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
    instance = new PolyaxonSdk.ProjectSearchesV1Api();
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

  describe('ProjectSearchesV1Api', function() {
    describe('createProjectSearch', function() {
      it('should call createProjectSearch successfully', function(done) {
        //uncomment below and update the code to test createProjectSearch
        //instance.createProjectSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('deleteProjectSearch', function() {
      it('should call deleteProjectSearch successfully', function(done) {
        //uncomment below and update the code to test deleteProjectSearch
        //instance.deleteProjectSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('getProjectSearch', function() {
      it('should call getProjectSearch successfully', function(done) {
        //uncomment below and update the code to test getProjectSearch
        //instance.getProjectSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listProjectSearchNames', function() {
      it('should call listProjectSearchNames successfully', function(done) {
        //uncomment below and update the code to test listProjectSearchNames
        //instance.listProjectSearchNames(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('listProjectSearches', function() {
      it('should call listProjectSearches successfully', function(done) {
        //uncomment below and update the code to test listProjectSearches
        //instance.listProjectSearches(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('patchProjectSearch', function() {
      it('should call patchProjectSearch successfully', function(done) {
        //uncomment below and update the code to test patchProjectSearch
        //instance.patchProjectSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('promoteProjectSearch', function() {
      it('should call promoteProjectSearch successfully', function(done) {
        //uncomment below and update the code to test promoteProjectSearch
        //instance.promoteProjectSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
    describe('updateProjectSearch', function() {
      it('should call updateProjectSearch successfully', function(done) {
        //uncomment below and update the code to test updateProjectSearch
        //instance.updateProjectSearch(function(error) {
        //  if (error) throw error;
        //expect().to.be();
        //});
        done();
      });
    });
  });

}));
