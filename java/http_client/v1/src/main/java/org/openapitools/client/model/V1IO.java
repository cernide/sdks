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
 * The version of the OpenAPI document: 1.13.0-rc2
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
 * V1IO
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1IO {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_DESCRIPTION = "description";
  @SerializedName(SERIALIZED_NAME_DESCRIPTION)
  private String description;

  public static final String SERIALIZED_NAME_TYPE = "type";
  @SerializedName(SERIALIZED_NAME_TYPE)
  private String type;

  public static final String SERIALIZED_NAME_VALUE = "value";
  @SerializedName(SERIALIZED_NAME_VALUE)
  private Object value;

  public static final String SERIALIZED_NAME_IS_OPTIONAL = "isOptional";
  @SerializedName(SERIALIZED_NAME_IS_OPTIONAL)
  private Boolean isOptional;

  public static final String SERIALIZED_NAME_IS_LIST = "isList";
  @SerializedName(SERIALIZED_NAME_IS_LIST)
  private Boolean isList;

  public static final String SERIALIZED_NAME_IS_FLAG = "isFlag";
  @SerializedName(SERIALIZED_NAME_IS_FLAG)
  private Boolean isFlag;

  public static final String SERIALIZED_NAME_ARG_FORMAT = "argFormat";
  @SerializedName(SERIALIZED_NAME_ARG_FORMAT)
  private String argFormat;

  public static final String SERIALIZED_NAME_DELAY_VALIDATION = "delayValidation";
  @SerializedName(SERIALIZED_NAME_DELAY_VALIDATION)
  private Boolean delayValidation;

  public static final String SERIALIZED_NAME_OPTIONS = "options";
  @SerializedName(SERIALIZED_NAME_OPTIONS)
  private List<Object> options = null;

  public static final String SERIALIZED_NAME_CONNECTION = "connection";
  @SerializedName(SERIALIZED_NAME_CONNECTION)
  private String connection;

  public static final String SERIALIZED_NAME_TO_INIT = "toInit";
  @SerializedName(SERIALIZED_NAME_TO_INIT)
  private Boolean toInit;

  public static final String SERIALIZED_NAME_TO_ENV = "toEnv";
  @SerializedName(SERIALIZED_NAME_TO_ENV)
  private String toEnv;


  public V1IO name(String name) {
    
    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public V1IO description(String description) {
    
    this.description = description;
    return this;
  }

   /**
   * Get description
   * @return description
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getDescription() {
    return description;
  }


  public void setDescription(String description) {
    this.description = description;
  }


  public V1IO type(String type) {
    
    this.type = type;
    return this;
  }

   /**
   * Get type
   * @return type
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getType() {
    return type;
  }


  public void setType(String type) {
    this.type = type;
  }


  public V1IO value(Object value) {
    
    this.value = value;
    return this;
  }

   /**
   * Get value
   * @return value
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getValue() {
    return value;
  }


  public void setValue(Object value) {
    this.value = value;
  }


  public V1IO isOptional(Boolean isOptional) {
    
    this.isOptional = isOptional;
    return this;
  }

   /**
   * Get isOptional
   * @return isOptional
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsOptional() {
    return isOptional;
  }


  public void setIsOptional(Boolean isOptional) {
    this.isOptional = isOptional;
  }


  public V1IO isList(Boolean isList) {
    
    this.isList = isList;
    return this;
  }

   /**
   * Get isList
   * @return isList
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsList() {
    return isList;
  }


  public void setIsList(Boolean isList) {
    this.isList = isList;
  }


  public V1IO isFlag(Boolean isFlag) {
    
    this.isFlag = isFlag;
    return this;
  }

   /**
   * Get isFlag
   * @return isFlag
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getIsFlag() {
    return isFlag;
  }


  public void setIsFlag(Boolean isFlag) {
    this.isFlag = isFlag;
  }


  public V1IO argFormat(String argFormat) {
    
    this.argFormat = argFormat;
    return this;
  }

   /**
   * Get argFormat
   * @return argFormat
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getArgFormat() {
    return argFormat;
  }


  public void setArgFormat(String argFormat) {
    this.argFormat = argFormat;
  }


  public V1IO delayValidation(Boolean delayValidation) {
    
    this.delayValidation = delayValidation;
    return this;
  }

   /**
   * Get delayValidation
   * @return delayValidation
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getDelayValidation() {
    return delayValidation;
  }


  public void setDelayValidation(Boolean delayValidation) {
    this.delayValidation = delayValidation;
  }


  public V1IO options(List<Object> options) {
    
    this.options = options;
    return this;
  }

  public V1IO addOptionsItem(Object optionsItem) {
    if (this.options == null) {
      this.options = new ArrayList<Object>();
    }
    this.options.add(optionsItem);
    return this;
  }

   /**
   * Get options
   * @return options
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getOptions() {
    return options;
  }


  public void setOptions(List<Object> options) {
    this.options = options;
  }


  public V1IO connection(String connection) {
    
    this.connection = connection;
    return this;
  }

   /**
   * Get connection
   * @return connection
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getConnection() {
    return connection;
  }


  public void setConnection(String connection) {
    this.connection = connection;
  }


  public V1IO toInit(Boolean toInit) {
    
    this.toInit = toInit;
    return this;
  }

   /**
   * Get toInit
   * @return toInit
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getToInit() {
    return toInit;
  }


  public void setToInit(Boolean toInit) {
    this.toInit = toInit;
  }


  public V1IO toEnv(String toEnv) {
    
    this.toEnv = toEnv;
    return this;
  }

   /**
   * Get toEnv
   * @return toEnv
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getToEnv() {
    return toEnv;
  }


  public void setToEnv(String toEnv) {
    this.toEnv = toEnv;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1IO v1IO = (V1IO) o;
    return Objects.equals(this.name, v1IO.name) &&
        Objects.equals(this.description, v1IO.description) &&
        Objects.equals(this.type, v1IO.type) &&
        Objects.equals(this.value, v1IO.value) &&
        Objects.equals(this.isOptional, v1IO.isOptional) &&
        Objects.equals(this.isList, v1IO.isList) &&
        Objects.equals(this.isFlag, v1IO.isFlag) &&
        Objects.equals(this.argFormat, v1IO.argFormat) &&
        Objects.equals(this.delayValidation, v1IO.delayValidation) &&
        Objects.equals(this.options, v1IO.options) &&
        Objects.equals(this.connection, v1IO.connection) &&
        Objects.equals(this.toInit, v1IO.toInit) &&
        Objects.equals(this.toEnv, v1IO.toEnv);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, description, type, value, isOptional, isList, isFlag, argFormat, delayValidation, options, connection, toInit, toEnv);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1IO {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    description: ").append(toIndentedString(description)).append("\n");
    sb.append("    type: ").append(toIndentedString(type)).append("\n");
    sb.append("    value: ").append(toIndentedString(value)).append("\n");
    sb.append("    isOptional: ").append(toIndentedString(isOptional)).append("\n");
    sb.append("    isList: ").append(toIndentedString(isList)).append("\n");
    sb.append("    isFlag: ").append(toIndentedString(isFlag)).append("\n");
    sb.append("    argFormat: ").append(toIndentedString(argFormat)).append("\n");
    sb.append("    delayValidation: ").append(toIndentedString(delayValidation)).append("\n");
    sb.append("    options: ").append(toIndentedString(options)).append("\n");
    sb.append("    connection: ").append(toIndentedString(connection)).append("\n");
    sb.append("    toInit: ").append(toIndentedString(toInit)).append("\n");
    sb.append("    toEnv: ").append(toIndentedString(toEnv)).append("\n");
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

