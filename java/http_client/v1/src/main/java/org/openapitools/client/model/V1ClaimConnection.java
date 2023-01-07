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
 * The version of the OpenAPI document: 1.21.0
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

/**
 * V1ClaimConnection
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1ClaimConnection {
  public static final String SERIALIZED_NAME_VOLUME_CLAIM = "volumeClaim";
  @SerializedName(SERIALIZED_NAME_VOLUME_CLAIM)
  private String volumeClaim;

  public static final String SERIALIZED_NAME_MOUNT_PATH = "mountPath";
  @SerializedName(SERIALIZED_NAME_MOUNT_PATH)
  private String mountPath;

  public static final String SERIALIZED_NAME_READ_ONLY = "readOnly";
  @SerializedName(SERIALIZED_NAME_READ_ONLY)
  private Boolean readOnly;

  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private Object kind;


  public V1ClaimConnection volumeClaim(String volumeClaim) {

    this.volumeClaim = volumeClaim;
    return this;
  }

   /**
   * Get volumeClaim
   * @return volumeClaim
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getVolumeClaim() {
    return volumeClaim;
  }


  public void setVolumeClaim(String volumeClaim) {
    this.volumeClaim = volumeClaim;
  }


  public V1ClaimConnection mountPath(String mountPath) {

    this.mountPath = mountPath;
    return this;
  }

   /**
   * Get mountPath
   * @return mountPath
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getMountPath() {
    return mountPath;
  }


  public void setMountPath(String mountPath) {
    this.mountPath = mountPath;
  }


  public V1ClaimConnection readOnly(Boolean readOnly) {

    this.readOnly = readOnly;
    return this;
  }

   /**
   * Get readOnly
   * @return readOnly
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Boolean getReadOnly() {
    return readOnly;
  }


  public void setReadOnly(Boolean readOnly) {
    this.readOnly = readOnly;
  }


  public V1ClaimConnection kind(Object kind) {

    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getKind() {
    return kind;
  }


  public void setKind(Object kind) {
    this.kind = kind;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1ClaimConnection v1ClaimConnection = (V1ClaimConnection) o;
    return Objects.equals(this.volumeClaim, v1ClaimConnection.volumeClaim) &&
        Objects.equals(this.mountPath, v1ClaimConnection.mountPath) &&
        Objects.equals(this.readOnly, v1ClaimConnection.readOnly) &&
        Objects.equals(this.kind, v1ClaimConnection.kind);
  }

  @Override
  public int hashCode() {
    return Objects.hash(volumeClaim, mountPath, readOnly, kind);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1ClaimConnection {\n");
    sb.append("    volumeClaim: ").append(toIndentedString(volumeClaim)).append("\n");
    sb.append("    mountPath: ").append(toIndentedString(mountPath)).append("\n");
    sb.append("    readOnly: ").append(toIndentedString(readOnly)).append("\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
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

