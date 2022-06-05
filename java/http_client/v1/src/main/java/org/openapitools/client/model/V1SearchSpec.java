// Copyright 2018-2022 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.18.2
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
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import org.openapitools.client.model.V1AnalyticsSpec;
import org.openapitools.client.model.V1DashboardSpec;

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


  public V1SearchSpec query(String query) {
    
    this.query = query;
    return this;
  }

   /**
   * Get query
   * @return query
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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
  @ApiModelProperty(value = "")

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

}

