// Copyright 2018-2021 Polyaxon, Inc.
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
 * The version of the OpenAPI document: 1.17.1
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
 * V1Cache
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Cache {
  public static final String SERIALIZED_NAME_DISABLE = "disable";
  @SerializedName(SERIALIZED_NAME_DISABLE)
  private Boolean disable;

  public static final String SERIALIZED_NAME_TTL = "ttl";
  @SerializedName(SERIALIZED_NAME_TTL)
  private Integer ttl;

  public static final String SERIALIZED_NAME_IO = "io";
  @SerializedName(SERIALIZED_NAME_IO)
  private List<String> io = null;

  public static final String SERIALIZED_NAME_SECTIONS = "sections";
  @SerializedName(SERIALIZED_NAME_SECTIONS)
  private List<String> sections = null;


  public V1Cache disable(Boolean disable) {
    
    this.disable = disable;
    return this;
  }

   /**
   * Get disable
   * @return disable
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDisable() {
    return disable;
  }


  public void setDisable(Boolean disable) {
    this.disable = disable;
  }


  public V1Cache ttl(Integer ttl) {
    
    this.ttl = ttl;
    return this;
  }

   /**
   * Get ttl
   * @return ttl
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Integer getTtl() {
    return ttl;
  }


  public void setTtl(Integer ttl) {
    this.ttl = ttl;
  }


  public V1Cache io(List<String> io) {
    
    this.io = io;
    return this;
  }

  public V1Cache addIoItem(String ioItem) {
    if (this.io == null) {
      this.io = new ArrayList<String>();
    }
    this.io.add(ioItem);
    return this;
  }

   /**
   * Get io
   * @return io
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getIo() {
    return io;
  }


  public void setIo(List<String> io) {
    this.io = io;
  }


  public V1Cache sections(List<String> sections) {
    
    this.sections = sections;
    return this;
  }

  public V1Cache addSectionsItem(String sectionsItem) {
    if (this.sections == null) {
      this.sections = new ArrayList<String>();
    }
    this.sections.add(sectionsItem);
    return this;
  }

   /**
   * Get sections
   * @return sections
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getSections() {
    return sections;
  }


  public void setSections(List<String> sections) {
    this.sections = sections;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Cache v1Cache = (V1Cache) o;
    return Objects.equals(this.disable, v1Cache.disable) &&
        Objects.equals(this.ttl, v1Cache.ttl) &&
        Objects.equals(this.io, v1Cache.io) &&
        Objects.equals(this.sections, v1Cache.sections);
  }

  @Override
  public int hashCode() {
    return Objects.hash(disable, ttl, io, sections);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Cache {\n");
    sb.append("    disable: ").append(toIndentedString(disable)).append("\n");
    sb.append("    ttl: ").append(toIndentedString(ttl)).append("\n");
    sb.append("    io: ").append(toIndentedString(io)).append("\n");
    sb.append("    sections: ").append(toIndentedString(sections)).append("\n");
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

