// Copyright 2018-2023 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 2.0.0-rc12
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import java.io.IOException;
import org.openapitools.client.model.V1AnalyticsSpec;
import org.openapitools.client.model.V1DashboardSpec;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import org.openapitools.client.JSON;

/**
 * V1SearchSpec
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1SearchSpec {
  public static final String SERIALIZED_NAME_QUERY = "query";
  @SerializedName(SERIALIZED_NAME_QUERY)
  private String query;

  public static final String SERIALIZED_NAME_SORT = "sort";
  @SerializedName(SERIALIZED_NAME_SORT)
  private String sort;

  public static final String SERIALIZED_NAME_LIMIT = "limit";
  @SerializedName(SERIALIZED_NAME_LIMIT)
  private Integer limit;

  public static final String SERIALIZED_NAME_OFFSET = "offset";
  @SerializedName(SERIALIZED_NAME_OFFSET)
  private Integer offset;

  public static final String SERIALIZED_NAME_GROUPBY = "groupby";
  @SerializedName(SERIALIZED_NAME_GROUPBY)
  private String groupby;

  public static final String SERIALIZED_NAME_COLUMNS = "columns";
  @SerializedName(SERIALIZED_NAME_COLUMNS)
  private String columns;

  public static final String SERIALIZED_NAME_LAYOUT = "layout";
  @SerializedName(SERIALIZED_NAME_LAYOUT)
  private String layout;

  public static final String SERIALIZED_NAME_SECTIONS = "sections";
  @SerializedName(SERIALIZED_NAME_SECTIONS)
  private String sections;

  public static final String SERIALIZED_NAME_COMPARES = "compares";
  @SerializedName(SERIALIZED_NAME_COMPARES)
  private String compares;

  public static final String SERIALIZED_NAME_HEAT = "heat";
  @SerializedName(SERIALIZED_NAME_HEAT)
  private String heat;

  public static final String SERIALIZED_NAME_EVENTS = "events";
  @SerializedName(SERIALIZED_NAME_EVENTS)
  private V1DashboardSpec events;

  public static final String SERIALIZED_NAME_HISTOGRAMS = "histograms";
  @SerializedName(SERIALIZED_NAME_HISTOGRAMS)
  private Object histograms;

  public static final String SERIALIZED_NAME_TRENDS = "trends";
  @SerializedName(SERIALIZED_NAME_TRENDS)
  private Object trends;

  public static final String SERIALIZED_NAME_ANALYTICS = "analytics";
  @SerializedName(SERIALIZED_NAME_ANALYTICS)
  private V1AnalyticsSpec analytics;

  public V1SearchSpec() {
  }

  public V1SearchSpec query(String query) {

    this.query = query;
    return this;
  }

   /**
   * Get query
   * @return query
  **/
  @javax.annotation.Nullable

  public String getQuery() {
    return query;
  }


  public void setQuery(String query) {
    this.query = query;
  }


  public V1SearchSpec sort(String sort) {

    this.sort = sort;
    return this;
  }

   /**
   * Get sort
   * @return sort
  **/
  @javax.annotation.Nullable

  public String getSort() {
    return sort;
  }


  public void setSort(String sort) {
    this.sort = sort;
  }


  public V1SearchSpec limit(Integer limit) {

    this.limit = limit;
    return this;
  }

   /**
   * Get limit
   * @return limit
  **/
  @javax.annotation.Nullable

  public Integer getLimit() {
    return limit;
  }


  public void setLimit(Integer limit) {
    this.limit = limit;
  }


  public V1SearchSpec offset(Integer offset) {

    this.offset = offset;
    return this;
  }

   /**
   * Get offset
   * @return offset
  **/
  @javax.annotation.Nullable

  public Integer getOffset() {
    return offset;
  }


  public void setOffset(Integer offset) {
    this.offset = offset;
  }


  public V1SearchSpec groupby(String groupby) {

    this.groupby = groupby;
    return this;
  }

   /**
   * Get groupby
   * @return groupby
  **/
  @javax.annotation.Nullable

  public String getGroupby() {
    return groupby;
  }


  public void setGroupby(String groupby) {
    this.groupby = groupby;
  }


  public V1SearchSpec columns(String columns) {

    this.columns = columns;
    return this;
  }

   /**
   * Get columns
   * @return columns
  **/
  @javax.annotation.Nullable

  public String getColumns() {
    return columns;
  }


  public void setColumns(String columns) {
    this.columns = columns;
  }


  public V1SearchSpec layout(String layout) {

    this.layout = layout;
    return this;
  }

   /**
   * Get layout
   * @return layout
  **/
  @javax.annotation.Nullable

  public String getLayout() {
    return layout;
  }


  public void setLayout(String layout) {
    this.layout = layout;
  }


  public V1SearchSpec sections(String sections) {

    this.sections = sections;
    return this;
  }

   /**
   * Get sections
   * @return sections
  **/
  @javax.annotation.Nullable

  public String getSections() {
    return sections;
  }


  public void setSections(String sections) {
    this.sections = sections;
  }


  public V1SearchSpec compares(String compares) {

    this.compares = compares;
    return this;
  }

   /**
   * Get compares
   * @return compares
  **/
  @javax.annotation.Nullable

  public String getCompares() {
    return compares;
  }


  public void setCompares(String compares) {
    this.compares = compares;
  }


  public V1SearchSpec heat(String heat) {

    this.heat = heat;
    return this;
  }

   /**
   * Get heat
   * @return heat
  **/
  @javax.annotation.Nullable

  public String getHeat() {
    return heat;
  }


  public void setHeat(String heat) {
    this.heat = heat;
  }


  public V1SearchSpec events(V1DashboardSpec events) {

    this.events = events;
    return this;
  }

   /**
   * Get events
   * @return events
  **/
  @javax.annotation.Nullable

  public V1DashboardSpec getEvents() {
    return events;
  }


  public void setEvents(V1DashboardSpec events) {
    this.events = events;
  }


  public V1SearchSpec histograms(Object histograms) {

    this.histograms = histograms;
    return this;
  }

   /**
   * Get histograms
   * @return histograms
  **/
  @javax.annotation.Nullable

  public Object getHistograms() {
    return histograms;
  }


  public void setHistograms(Object histograms) {
    this.histograms = histograms;
  }


  public V1SearchSpec trends(Object trends) {

    this.trends = trends;
    return this;
  }

   /**
   * Get trends
   * @return trends
  **/
  @javax.annotation.Nullable

  public Object getTrends() {
    return trends;
  }


  public void setTrends(Object trends) {
    this.trends = trends;
  }


  public V1SearchSpec analytics(V1AnalyticsSpec analytics) {

    this.analytics = analytics;
    return this;
  }

   /**
   * Get analytics
   * @return analytics
  **/
  @javax.annotation.Nullable

  public V1AnalyticsSpec getAnalytics() {
    return analytics;
  }


  public void setAnalytics(V1AnalyticsSpec analytics) {
    this.analytics = analytics;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1SearchSpec v1SearchSpec = (V1SearchSpec) o;
    return Objects.equals(this.query, v1SearchSpec.query) &&
        Objects.equals(this.sort, v1SearchSpec.sort) &&
        Objects.equals(this.limit, v1SearchSpec.limit) &&
        Objects.equals(this.offset, v1SearchSpec.offset) &&
        Objects.equals(this.groupby, v1SearchSpec.groupby) &&
        Objects.equals(this.columns, v1SearchSpec.columns) &&
        Objects.equals(this.layout, v1SearchSpec.layout) &&
        Objects.equals(this.sections, v1SearchSpec.sections) &&
        Objects.equals(this.compares, v1SearchSpec.compares) &&
        Objects.equals(this.heat, v1SearchSpec.heat) &&
        Objects.equals(this.events, v1SearchSpec.events) &&
        Objects.equals(this.histograms, v1SearchSpec.histograms) &&
        Objects.equals(this.trends, v1SearchSpec.trends) &&
        Objects.equals(this.analytics, v1SearchSpec.analytics);
  }

  @Override
  public int hashCode() {
    return Objects.hash(query, sort, limit, offset, groupby, columns, layout, sections, compares, heat, events, histograms, trends, analytics);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1SearchSpec {\n");
    sb.append("    query: ").append(toIndentedString(query)).append("\n");
    sb.append("    sort: ").append(toIndentedString(sort)).append("\n");
    sb.append("    limit: ").append(toIndentedString(limit)).append("\n");
    sb.append("    offset: ").append(toIndentedString(offset)).append("\n");
    sb.append("    groupby: ").append(toIndentedString(groupby)).append("\n");
    sb.append("    columns: ").append(toIndentedString(columns)).append("\n");
    sb.append("    layout: ").append(toIndentedString(layout)).append("\n");
    sb.append("    sections: ").append(toIndentedString(sections)).append("\n");
    sb.append("    compares: ").append(toIndentedString(compares)).append("\n");
    sb.append("    heat: ").append(toIndentedString(heat)).append("\n");
    sb.append("    events: ").append(toIndentedString(events)).append("\n");
    sb.append("    histograms: ").append(toIndentedString(histograms)).append("\n");
    sb.append("    trends: ").append(toIndentedString(trends)).append("\n");
    sb.append("    analytics: ").append(toIndentedString(analytics)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("query");
    openapiFields.add("sort");
    openapiFields.add("limit");
    openapiFields.add("offset");
    openapiFields.add("groupby");
    openapiFields.add("columns");
    openapiFields.add("layout");
    openapiFields.add("sections");
    openapiFields.add("compares");
    openapiFields.add("heat");
    openapiFields.add("events");
    openapiFields.add("histograms");
    openapiFields.add("trends");
    openapiFields.add("analytics");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1SearchSpec
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1SearchSpec.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1SearchSpec is not found in the empty JSON string", V1SearchSpec.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1SearchSpec.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1SearchSpec` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("query") != null && !jsonObj.get("query").isJsonNull()) && !jsonObj.get("query").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `query` to be a primitive type in the JSON string but got `%s`", jsonObj.get("query").toString()));
      }
      if ((jsonObj.get("sort") != null && !jsonObj.get("sort").isJsonNull()) && !jsonObj.get("sort").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `sort` to be a primitive type in the JSON string but got `%s`", jsonObj.get("sort").toString()));
      }
      if ((jsonObj.get("groupby") != null && !jsonObj.get("groupby").isJsonNull()) && !jsonObj.get("groupby").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `groupby` to be a primitive type in the JSON string but got `%s`", jsonObj.get("groupby").toString()));
      }
      if ((jsonObj.get("columns") != null && !jsonObj.get("columns").isJsonNull()) && !jsonObj.get("columns").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `columns` to be a primitive type in the JSON string but got `%s`", jsonObj.get("columns").toString()));
      }
      if ((jsonObj.get("layout") != null && !jsonObj.get("layout").isJsonNull()) && !jsonObj.get("layout").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `layout` to be a primitive type in the JSON string but got `%s`", jsonObj.get("layout").toString()));
      }
      if ((jsonObj.get("sections") != null && !jsonObj.get("sections").isJsonNull()) && !jsonObj.get("sections").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `sections` to be a primitive type in the JSON string but got `%s`", jsonObj.get("sections").toString()));
      }
      if ((jsonObj.get("compares") != null && !jsonObj.get("compares").isJsonNull()) && !jsonObj.get("compares").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `compares` to be a primitive type in the JSON string but got `%s`", jsonObj.get("compares").toString()));
      }
      if ((jsonObj.get("heat") != null && !jsonObj.get("heat").isJsonNull()) && !jsonObj.get("heat").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `heat` to be a primitive type in the JSON string but got `%s`", jsonObj.get("heat").toString()));
      }
      // validate the optional field `events`
      if (jsonObj.get("events") != null && !jsonObj.get("events").isJsonNull()) {
        V1DashboardSpec.validateJsonObject(jsonObj.getAsJsonObject("events"));
      }
      // validate the optional field `analytics`
      if (jsonObj.get("analytics") != null && !jsonObj.get("analytics").isJsonNull()) {
        V1AnalyticsSpec.validateJsonObject(jsonObj.getAsJsonObject("analytics"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1SearchSpec.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1SearchSpec' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1SearchSpec> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1SearchSpec.class));

       return (TypeAdapter<T>) new TypeAdapter<V1SearchSpec>() {
           @Override
           public void write(JsonWriter out, V1SearchSpec value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1SearchSpec read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1SearchSpec given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1SearchSpec
  * @throws IOException if the JSON string is invalid with respect to V1SearchSpec
  */
  public static V1SearchSpec fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1SearchSpec.class);
  }

 /**
  * Convert an instance of V1SearchSpec to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

