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
 * The version of the OpenAPI document: 1.22.0-rc0
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
import java.util.ArrayList;
import java.util.List;

/**
 * V1TensorboardType
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1TensorboardType {
  public static final String SERIALIZED_NAME_PORT = "port";
  @SerializedName(SERIALIZED_NAME_PORT)
  private Integer port;

  public static final String SERIALIZED_NAME_UUIDS = "uuids";
  @SerializedName(SERIALIZED_NAME_UUIDS)
  private List<String> uuids = null;

  public static final String SERIALIZED_NAME_USE_NAMES = "use_names";
  @SerializedName(SERIALIZED_NAME_USE_NAMES)
  private Boolean useNames;

  public static final String SERIALIZED_NAME_PATH_PREFIX = "path_prefix";
  @SerializedName(SERIALIZED_NAME_PATH_PREFIX)
  private String pathPrefix;

  public static final String SERIALIZED_NAME_PLUGINS = "plugins";
  @SerializedName(SERIALIZED_NAME_PLUGINS)
  private String plugins;


  public V1TensorboardType port(Integer port) {
    
    this.port = port;
    return this;
  }

   /**
   * Get port
   * @return port
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getPort() {
    return port;
  }


  public void setPort(Integer port) {
    this.port = port;
  }


  public V1TensorboardType uuids(List<String> uuids) {
    
    this.uuids = uuids;
    return this;
  }

  public V1TensorboardType addUuidsItem(String uuidsItem) {
    if (this.uuids == null) {
      this.uuids = new ArrayList<String>();
    }
    this.uuids.add(uuidsItem);
    return this;
  }

   /**
   * Get uuids
   * @return uuids
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getUuids() {
    return uuids;
  }


  public void setUuids(List<String> uuids) {
    this.uuids = uuids;
  }


  public V1TensorboardType useNames(Boolean useNames) {
    
    this.useNames = useNames;
    return this;
  }

   /**
   * Get useNames
   * @return useNames
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getUseNames() {
    return useNames;
  }


  public void setUseNames(Boolean useNames) {
    this.useNames = useNames;
  }


  public V1TensorboardType pathPrefix(String pathPrefix) {
    
    this.pathPrefix = pathPrefix;
    return this;
  }

   /**
   * Get pathPrefix
   * @return pathPrefix
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPathPrefix() {
    return pathPrefix;
  }


  public void setPathPrefix(String pathPrefix) {
    this.pathPrefix = pathPrefix;
  }


  public V1TensorboardType plugins(String plugins) {
    
    this.plugins = plugins;
    return this;
  }

   /**
   * Get plugins
   * @return plugins
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPlugins() {
    return plugins;
  }


  public void setPlugins(String plugins) {
    this.plugins = plugins;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1TensorboardType v1TensorboardType = (V1TensorboardType) o;
    return Objects.equals(this.port, v1TensorboardType.port) &&
        Objects.equals(this.uuids, v1TensorboardType.uuids) &&
        Objects.equals(this.useNames, v1TensorboardType.useNames) &&
        Objects.equals(this.pathPrefix, v1TensorboardType.pathPrefix) &&
        Objects.equals(this.plugins, v1TensorboardType.plugins);
  }

  @Override
  public int hashCode() {
    return Objects.hash(port, uuids, useNames, pathPrefix, plugins);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1TensorboardType {\n");
    sb.append("    port: ").append(toIndentedString(port)).append("\n");
    sb.append("    uuids: ").append(toIndentedString(uuids)).append("\n");
    sb.append("    useNames: ").append(toIndentedString(useNames)).append("\n");
    sb.append("    pathPrefix: ").append(toIndentedString(pathPrefix)).append("\n");
    sb.append("    plugins: ").append(toIndentedString(plugins)).append("\n");
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

