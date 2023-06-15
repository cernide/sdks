/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc19
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.RuntimeError;
import org.openapitools.client.model.V1EntitiesTags;
import org.openapitools.client.model.V1ListTagsResponse;
import org.openapitools.client.model.V1Tag;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for TagsV1Api
 */
@Disabled
public class TagsV1ApiTest {

    private final TagsV1Api api = new TagsV1Api();

    /**
     * Create tag
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void createTagTest() throws ApiException {
        String owner = null;
        V1Tag body = null;
        V1Tag response = api.createTag(owner, body);
        // TODO: test validations
    }

    /**
     * Delete tag
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void deleteTagTest() throws ApiException {
        String owner = null;
        String uuid = null;
        Boolean cascade = null;
        api.deleteTag(owner, uuid, cascade);
        // TODO: test validations
    }

    /**
     * Get tag
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getTagTest() throws ApiException {
        String owner = null;
        String uuid = null;
        V1Tag response = api.getTag(owner, uuid);
        // TODO: test validations
    }

    /**
     * List tags
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void listTagsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        V1ListTagsResponse response = api.listTags(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Load tags
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void loadTagsTest() throws ApiException {
        String owner = null;
        Integer offset = null;
        Integer limit = null;
        String sort = null;
        String query = null;
        Boolean bookmarks = null;
        String mode = null;
        Boolean noPage = null;
        Object response = api.loadTags(owner, offset, limit, sort, query, bookmarks, mode, noPage);
        // TODO: test validations
    }

    /**
     * Patch tag
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void patchTagTest() throws ApiException {
        String owner = null;
        String tagUuid = null;
        V1Tag body = null;
        V1Tag response = api.patchTag(owner, tagUuid, body);
        // TODO: test validations
    }

    /**
     * Sync tags
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void syncTagsTest() throws ApiException {
        String owner = null;
        V1EntitiesTags body = null;
        api.syncTags(owner, body);
        // TODO: test validations
    }

    /**
     * Update tag
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void updateTagTest() throws ApiException {
        String owner = null;
        String tagUuid = null;
        V1Tag body = null;
        V1Tag response = api.updateTag(owner, tagUuid, body);
        // TODO: test validations
    }

}
