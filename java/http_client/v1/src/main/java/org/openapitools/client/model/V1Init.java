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
import org.openapitools.client.model.V1ArtifactsType;
import org.openapitools.client.model.V1DockerfileType;
import org.openapitools.client.model.V1FileType;
import org.openapitools.client.model.V1GitType;
import org.openapitools.client.model.V1TensorboardType;

/**
 * V1Init
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Init {
  public static final String SERIALIZED_NAME_ARTIFACTS = "artifacts";
  @SerializedName(SERIALIZED_NAME_ARTIFACTS)
  private V1ArtifactsType artifacts;

  public static final String SERIALIZED_NAME_PATHS = "paths";
  @SerializedName(SERIALIZED_NAME_PATHS)
  private List<Object> paths = null;

  public static final String SERIALIZED_NAME_GIT = "git";
  @SerializedName(SERIALIZED_NAME_GIT)
  private V1GitType git;

  public static final String SERIALIZED_NAME_DOCKERFILE = "dockerfile";
  @SerializedName(SERIALIZED_NAME_DOCKERFILE)
  private V1DockerfileType dockerfile;

  public static final String SERIALIZED_NAME_FILE = "file";
  @SerializedName(SERIALIZED_NAME_FILE)
  private V1FileType file;

  public static final String SERIALIZED_NAME_TENSORBOARD = "tensorboard";
  @SerializedName(SERIALIZED_NAME_TENSORBOARD)
  private V1TensorboardType tensorboard;

  public static final String SERIALIZED_NAME_LINEAGE_REF = "lineageRef";
  @SerializedName(SERIALIZED_NAME_LINEAGE_REF)
  private String lineageRef;

  public static final String SERIALIZED_NAME_ARTIFACT_REF = "artifactRef";
  @SerializedName(SERIALIZED_NAME_ARTIFACT_REF)
  private String artifactRef;

  public static final String SERIALIZED_NAME_MODEL_REF = "modelRef";
  @SerializedName(SERIALIZED_NAME_MODEL_REF)
  private String modelRef;

  public static final String SERIALIZED_NAME_CONNECTION = "connection";
  @SerializedName(SERIALIZED_NAME_CONNECTION)
  private String connection;

  public static final String SERIALIZED_NAME_PATH = "path";
  @SerializedName(SERIALIZED_NAME_PATH)
  private String path;

  public static final String SERIALIZED_NAME_CONTAINER = "container";
  @SerializedName(SERIALIZED_NAME_CONTAINER)
  private Object container;


  public V1Init artifacts(V1ArtifactsType artifacts) {
    
    this.artifacts = artifacts;
    return this;
  }

   /**
   * Get artifacts
   * @return artifacts
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1ArtifactsType getArtifacts() {
    return artifacts;
  }


  public void setArtifacts(V1ArtifactsType artifacts) {
    this.artifacts = artifacts;
  }


  public V1Init paths(List<Object> paths) {
    
    this.paths = paths;
    return this;
  }

  public V1Init addPathsItem(Object pathsItem) {
    if (this.paths == null) {
      this.paths = new ArrayList<Object>();
    }
    this.paths.add(pathsItem);
    return this;
  }

   /**
   * Get paths
   * @return paths
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<Object> getPaths() {
    return paths;
  }


  public void setPaths(List<Object> paths) {
    this.paths = paths;
  }


  public V1Init git(V1GitType git) {
    
    this.git = git;
    return this;
  }

   /**
   * Get git
   * @return git
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1GitType getGit() {
    return git;
  }


  public void setGit(V1GitType git) {
    this.git = git;
  }


  public V1Init dockerfile(V1DockerfileType dockerfile) {
    
    this.dockerfile = dockerfile;
    return this;
  }

   /**
   * Get dockerfile
   * @return dockerfile
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1DockerfileType getDockerfile() {
    return dockerfile;
  }


  public void setDockerfile(V1DockerfileType dockerfile) {
    this.dockerfile = dockerfile;
  }


  public V1Init file(V1FileType file) {
    
    this.file = file;
    return this;
  }

   /**
   * Get file
   * @return file
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1FileType getFile() {
    return file;
  }


  public void setFile(V1FileType file) {
    this.file = file;
  }


  public V1Init tensorboard(V1TensorboardType tensorboard) {
    
    this.tensorboard = tensorboard;
    return this;
  }

   /**
   * Get tensorboard
   * @return tensorboard
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1TensorboardType getTensorboard() {
    return tensorboard;
  }


  public void setTensorboard(V1TensorboardType tensorboard) {
    this.tensorboard = tensorboard;
  }


  public V1Init lineageRef(String lineageRef) {
    
    this.lineageRef = lineageRef;
    return this;
  }

   /**
   * Get lineageRef
   * @return lineageRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getLineageRef() {
    return lineageRef;
  }


  public void setLineageRef(String lineageRef) {
    this.lineageRef = lineageRef;
  }


  public V1Init artifactRef(String artifactRef) {
    
    this.artifactRef = artifactRef;
    return this;
  }

   /**
   * Get artifactRef
   * @return artifactRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getArtifactRef() {
    return artifactRef;
  }


  public void setArtifactRef(String artifactRef) {
    this.artifactRef = artifactRef;
  }


  public V1Init modelRef(String modelRef) {
    
    this.modelRef = modelRef;
    return this;
  }

   /**
   * Get modelRef
   * @return modelRef
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getModelRef() {
    return modelRef;
  }


  public void setModelRef(String modelRef) {
    this.modelRef = modelRef;
  }


  public V1Init connection(String connection) {
    
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


  public V1Init path(String path) {
    
    this.path = path;
    return this;
  }

   /**
   * Get path
   * @return path
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getPath() {
    return path;
  }


  public void setPath(String path) {
    this.path = path;
  }


  public V1Init container(Object container) {
    
    this.container = container;
    return this;
  }

   /**
   * Get container
   * @return container
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public Object getContainer() {
    return container;
  }


  public void setContainer(Object container) {
    this.container = container;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Init v1Init = (V1Init) o;
    return Objects.equals(this.artifacts, v1Init.artifacts) &&
        Objects.equals(this.paths, v1Init.paths) &&
        Objects.equals(this.git, v1Init.git) &&
        Objects.equals(this.dockerfile, v1Init.dockerfile) &&
        Objects.equals(this.file, v1Init.file) &&
        Objects.equals(this.tensorboard, v1Init.tensorboard) &&
        Objects.equals(this.lineageRef, v1Init.lineageRef) &&
        Objects.equals(this.artifactRef, v1Init.artifactRef) &&
        Objects.equals(this.modelRef, v1Init.modelRef) &&
        Objects.equals(this.connection, v1Init.connection) &&
        Objects.equals(this.path, v1Init.path) &&
        Objects.equals(this.container, v1Init.container);
  }

  @Override
  public int hashCode() {
    return Objects.hash(artifacts, paths, git, dockerfile, file, tensorboard, lineageRef, artifactRef, modelRef, connection, path, container);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Init {\n");
    sb.append("    artifacts: ").append(toIndentedString(artifacts)).append("\n");
    sb.append("    paths: ").append(toIndentedString(paths)).append("\n");
    sb.append("    git: ").append(toIndentedString(git)).append("\n");
    sb.append("    dockerfile: ").append(toIndentedString(dockerfile)).append("\n");
    sb.append("    file: ").append(toIndentedString(file)).append("\n");
    sb.append("    tensorboard: ").append(toIndentedString(tensorboard)).append("\n");
    sb.append("    lineageRef: ").append(toIndentedString(lineageRef)).append("\n");
    sb.append("    artifactRef: ").append(toIndentedString(artifactRef)).append("\n");
    sb.append("    modelRef: ").append(toIndentedString(modelRef)).append("\n");
    sb.append("    connection: ").append(toIndentedString(connection)).append("\n");
    sb.append("    path: ").append(toIndentedString(path)).append("\n");
    sb.append("    container: ").append(toIndentedString(container)).append("\n");
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

