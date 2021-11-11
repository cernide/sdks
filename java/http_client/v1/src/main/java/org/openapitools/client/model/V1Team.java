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
 * The version of the OpenAPI document: 1.13.0-rc1
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
import org.openapitools.client.model.V1TeamSettings;
import org.threeten.bp.OffsetDateTime;

/**
 * V1Team
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Team {
  public static final String SERIALIZED_NAME_UUID = "uuid";
  @SerializedName(SERIALIZED_NAME_UUID)
  private String uuid;

  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_PROJECTS = "projects";
  @SerializedName(SERIALIZED_NAME_PROJECTS)
  private List<String> projects = null;

  public static final String SERIALIZED_NAME_COMPONENT_HUBS = "component_hubs";
  @SerializedName(SERIALIZED_NAME_COMPONENT_HUBS)
  private List<String> componentHubs = null;

  public static final String SERIALIZED_NAME_MODEL_REGISTRIES = "model_registries";
  @SerializedName(SERIALIZED_NAME_MODEL_REGISTRIES)
  private List<String> modelRegistries = null;

  public static final String SERIALIZED_NAME_SETTINGS = "settings";
  @SerializedName(SERIALIZED_NAME_SETTINGS)
  private V1TeamSettings settings;

  public static final String SERIALIZED_NAME_CREATED_AT = "created_at";
  @SerializedName(SERIALIZED_NAME_CREATED_AT)
  private OffsetDateTime createdAt;

  public static final String SERIALIZED_NAME_UPDATED_AT = "updated_at";
  @SerializedName(SERIALIZED_NAME_UPDATED_AT)
  private OffsetDateTime updatedAt;


  public V1Team uuid(String uuid) {
    
    this.uuid = uuid;
    return this;
  }

   /**
   * Get uuid
   * @return uuid
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public String getUuid() {
    return uuid;
  }


  public void setUuid(String uuid) {
    this.uuid = uuid;
  }


  public V1Team name(String name) {
    
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


  public V1Team projects(List<String> projects) {
    
    this.projects = projects;
    return this;
  }

  public V1Team addProjectsItem(String projectsItem) {
    if (this.projects == null) {
      this.projects = new ArrayList<String>();
    }
    this.projects.add(projectsItem);
    return this;
  }

   /**
   * Get projects
   * @return projects
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getProjects() {
    return projects;
  }


  public void setProjects(List<String> projects) {
    this.projects = projects;
  }


  public V1Team componentHubs(List<String> componentHubs) {
    
    this.componentHubs = componentHubs;
    return this;
  }

  public V1Team addComponentHubsItem(String componentHubsItem) {
    if (this.componentHubs == null) {
      this.componentHubs = new ArrayList<String>();
    }
    this.componentHubs.add(componentHubsItem);
    return this;
  }

   /**
   * Get componentHubs
   * @return componentHubs
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getComponentHubs() {
    return componentHubs;
  }


  public void setComponentHubs(List<String> componentHubs) {
    this.componentHubs = componentHubs;
  }


  public V1Team modelRegistries(List<String> modelRegistries) {
    
    this.modelRegistries = modelRegistries;
    return this;
  }

  public V1Team addModelRegistriesItem(String modelRegistriesItem) {
    if (this.modelRegistries == null) {
      this.modelRegistries = new ArrayList<String>();
    }
    this.modelRegistries.add(modelRegistriesItem);
    return this;
  }

   /**
   * Get modelRegistries
   * @return modelRegistries
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public List<String> getModelRegistries() {
    return modelRegistries;
  }


  public void setModelRegistries(List<String> modelRegistries) {
    this.modelRegistries = modelRegistries;
  }


  public V1Team settings(V1TeamSettings settings) {
    
    this.settings = settings;
    return this;
  }

   /**
   * Get settings
   * @return settings
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public V1TeamSettings getSettings() {
    return settings;
  }


  public void setSettings(V1TeamSettings settings) {
    this.settings = settings;
  }


  public V1Team createdAt(OffsetDateTime createdAt) {
    
    this.createdAt = createdAt;
    return this;
  }

   /**
   * Get createdAt
   * @return createdAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getCreatedAt() {
    return createdAt;
  }


  public void setCreatedAt(OffsetDateTime createdAt) {
    this.createdAt = createdAt;
  }


  public V1Team updatedAt(OffsetDateTime updatedAt) {
    
    this.updatedAt = updatedAt;
    return this;
  }

   /**
   * Get updatedAt
   * @return updatedAt
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "")

  public OffsetDateTime getUpdatedAt() {
    return updatedAt;
  }


  public void setUpdatedAt(OffsetDateTime updatedAt) {
    this.updatedAt = updatedAt;
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Team v1Team = (V1Team) o;
    return Objects.equals(this.uuid, v1Team.uuid) &&
        Objects.equals(this.name, v1Team.name) &&
        Objects.equals(this.projects, v1Team.projects) &&
        Objects.equals(this.componentHubs, v1Team.componentHubs) &&
        Objects.equals(this.modelRegistries, v1Team.modelRegistries) &&
        Objects.equals(this.settings, v1Team.settings) &&
        Objects.equals(this.createdAt, v1Team.createdAt) &&
        Objects.equals(this.updatedAt, v1Team.updatedAt);
  }

  @Override
  public int hashCode() {
    return Objects.hash(uuid, name, projects, componentHubs, modelRegistries, settings, createdAt, updatedAt);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Team {\n");
    sb.append("    uuid: ").append(toIndentedString(uuid)).append("\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    projects: ").append(toIndentedString(projects)).append("\n");
    sb.append("    componentHubs: ").append(toIndentedString(componentHubs)).append("\n");
    sb.append("    modelRegistries: ").append(toIndentedString(modelRegistries)).append("\n");
    sb.append("    settings: ").append(toIndentedString(settings)).append("\n");
    sb.append("    createdAt: ").append(toIndentedString(createdAt)).append("\n");
    sb.append("    updatedAt: ").append(toIndentedString(updatedAt)).append("\n");
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

