/*
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.6-rc0
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
import java.util.ArrayList;
import java.util.List;
import org.openapitools.client.model.V1UserAccess;

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
 * V1ProjectSettings
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1ProjectSettings {
  public static final String SERIALIZED_NAME_CONNECTIONS = "connections";
  @SerializedName(SERIALIZED_NAME_CONNECTIONS)
  private List<String> connections;

  public static final String SERIALIZED_NAME_PRESET = "preset";
  @SerializedName(SERIALIZED_NAME_PRESET)
  private String preset;

  public static final String SERIALIZED_NAME_PRESETS = "presets";
  @SerializedName(SERIALIZED_NAME_PRESETS)
  private List<String> presets;

  public static final String SERIALIZED_NAME_QUEUE = "queue";
  @SerializedName(SERIALIZED_NAME_QUEUE)
  private String queue;

  public static final String SERIALIZED_NAME_QUEUES = "queues";
  @SerializedName(SERIALIZED_NAME_QUEUES)
  private List<String> queues;

  public static final String SERIALIZED_NAME_AGENTS = "agents";
  @SerializedName(SERIALIZED_NAME_AGENTS)
  private List<String> agents;

  public static final String SERIALIZED_NAME_NAMESPACES = "namespaces";
  @SerializedName(SERIALIZED_NAME_NAMESPACES)
  private List<String> namespaces;

  public static final String SERIALIZED_NAME_USER_ACCESSES = "user_accesses";
  @SerializedName(SERIALIZED_NAME_USER_ACCESSES)
  private List<V1UserAccess> userAccesses;

  public static final String SERIALIZED_NAME_TEAMS = "teams";
  @SerializedName(SERIALIZED_NAME_TEAMS)
  private List<String> teams;

  public static final String SERIALIZED_NAME_PROJECTS = "projects";
  @SerializedName(SERIALIZED_NAME_PROJECTS)
  private List<String> projects;

  public static final String SERIALIZED_NAME_POLICY = "policy";
  @SerializedName(SERIALIZED_NAME_POLICY)
  private String policy;

  public V1ProjectSettings() {
  }

  public V1ProjectSettings connections(List<String> connections) {
    
    this.connections = connections;
    return this;
  }

  public V1ProjectSettings addConnectionsItem(String connectionsItem) {
    if (this.connections == null) {
      this.connections = new ArrayList<>();
    }
    this.connections.add(connectionsItem);
    return this;
  }

   /**
   * Get connections
   * @return connections
  **/
  @javax.annotation.Nullable

  public List<String> getConnections() {
    return connections;
  }


  public void setConnections(List<String> connections) {
    this.connections = connections;
  }


  public V1ProjectSettings preset(String preset) {
    
    this.preset = preset;
    return this;
  }

   /**
   * Get preset
   * @return preset
  **/
  @javax.annotation.Nullable

  public String getPreset() {
    return preset;
  }


  public void setPreset(String preset) {
    this.preset = preset;
  }


  public V1ProjectSettings presets(List<String> presets) {
    
    this.presets = presets;
    return this;
  }

  public V1ProjectSettings addPresetsItem(String presetsItem) {
    if (this.presets == null) {
      this.presets = new ArrayList<>();
    }
    this.presets.add(presetsItem);
    return this;
  }

   /**
   * Get presets
   * @return presets
  **/
  @javax.annotation.Nullable

  public List<String> getPresets() {
    return presets;
  }


  public void setPresets(List<String> presets) {
    this.presets = presets;
  }


  public V1ProjectSettings queue(String queue) {
    
    this.queue = queue;
    return this;
  }

   /**
   * Get queue
   * @return queue
  **/
  @javax.annotation.Nullable

  public String getQueue() {
    return queue;
  }


  public void setQueue(String queue) {
    this.queue = queue;
  }


  public V1ProjectSettings queues(List<String> queues) {
    
    this.queues = queues;
    return this;
  }

  public V1ProjectSettings addQueuesItem(String queuesItem) {
    if (this.queues == null) {
      this.queues = new ArrayList<>();
    }
    this.queues.add(queuesItem);
    return this;
  }

   /**
   * Get queues
   * @return queues
  **/
  @javax.annotation.Nullable

  public List<String> getQueues() {
    return queues;
  }


  public void setQueues(List<String> queues) {
    this.queues = queues;
  }


  public V1ProjectSettings agents(List<String> agents) {
    
    this.agents = agents;
    return this;
  }

  public V1ProjectSettings addAgentsItem(String agentsItem) {
    if (this.agents == null) {
      this.agents = new ArrayList<>();
    }
    this.agents.add(agentsItem);
    return this;
  }

   /**
   * Get agents
   * @return agents
  **/
  @javax.annotation.Nullable

  public List<String> getAgents() {
    return agents;
  }


  public void setAgents(List<String> agents) {
    this.agents = agents;
  }


  public V1ProjectSettings namespaces(List<String> namespaces) {
    
    this.namespaces = namespaces;
    return this;
  }

  public V1ProjectSettings addNamespacesItem(String namespacesItem) {
    if (this.namespaces == null) {
      this.namespaces = new ArrayList<>();
    }
    this.namespaces.add(namespacesItem);
    return this;
  }

   /**
   * Get namespaces
   * @return namespaces
  **/
  @javax.annotation.Nullable

  public List<String> getNamespaces() {
    return namespaces;
  }


  public void setNamespaces(List<String> namespaces) {
    this.namespaces = namespaces;
  }


  public V1ProjectSettings userAccesses(List<V1UserAccess> userAccesses) {
    
    this.userAccesses = userAccesses;
    return this;
  }

  public V1ProjectSettings addUserAccessesItem(V1UserAccess userAccessesItem) {
    if (this.userAccesses == null) {
      this.userAccesses = new ArrayList<>();
    }
    this.userAccesses.add(userAccessesItem);
    return this;
  }

   /**
   * Get userAccesses
   * @return userAccesses
  **/
  @javax.annotation.Nullable

  public List<V1UserAccess> getUserAccesses() {
    return userAccesses;
  }


  public void setUserAccesses(List<V1UserAccess> userAccesses) {
    this.userAccesses = userAccesses;
  }


  public V1ProjectSettings teams(List<String> teams) {
    
    this.teams = teams;
    return this;
  }

  public V1ProjectSettings addTeamsItem(String teamsItem) {
    if (this.teams == null) {
      this.teams = new ArrayList<>();
    }
    this.teams.add(teamsItem);
    return this;
  }

   /**
   * Get teams
   * @return teams
  **/
  @javax.annotation.Nullable

  public List<String> getTeams() {
    return teams;
  }


  public void setTeams(List<String> teams) {
    this.teams = teams;
  }


  public V1ProjectSettings projects(List<String> projects) {
    
    this.projects = projects;
    return this;
  }

  public V1ProjectSettings addProjectsItem(String projectsItem) {
    if (this.projects == null) {
      this.projects = new ArrayList<>();
    }
    this.projects.add(projectsItem);
    return this;
  }

   /**
   * Get projects
   * @return projects
  **/
  @javax.annotation.Nullable

  public List<String> getProjects() {
    return projects;
  }


  public void setProjects(List<String> projects) {
    this.projects = projects;
  }


  public V1ProjectSettings policy(String policy) {
    
    this.policy = policy;
    return this;
  }

   /**
   * Get policy
   * @return policy
  **/
  @javax.annotation.Nullable

  public String getPolicy() {
    return policy;
  }


  public void setPolicy(String policy) {
    this.policy = policy;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1ProjectSettings v1ProjectSettings = (V1ProjectSettings) o;
    return Objects.equals(this.connections, v1ProjectSettings.connections) &&
        Objects.equals(this.preset, v1ProjectSettings.preset) &&
        Objects.equals(this.presets, v1ProjectSettings.presets) &&
        Objects.equals(this.queue, v1ProjectSettings.queue) &&
        Objects.equals(this.queues, v1ProjectSettings.queues) &&
        Objects.equals(this.agents, v1ProjectSettings.agents) &&
        Objects.equals(this.namespaces, v1ProjectSettings.namespaces) &&
        Objects.equals(this.userAccesses, v1ProjectSettings.userAccesses) &&
        Objects.equals(this.teams, v1ProjectSettings.teams) &&
        Objects.equals(this.projects, v1ProjectSettings.projects) &&
        Objects.equals(this.policy, v1ProjectSettings.policy);
  }

  @Override
  public int hashCode() {
    return Objects.hash(connections, preset, presets, queue, queues, agents, namespaces, userAccesses, teams, projects, policy);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1ProjectSettings {\n");
    sb.append("    connections: ").append(toIndentedString(connections)).append("\n");
    sb.append("    preset: ").append(toIndentedString(preset)).append("\n");
    sb.append("    presets: ").append(toIndentedString(presets)).append("\n");
    sb.append("    queue: ").append(toIndentedString(queue)).append("\n");
    sb.append("    queues: ").append(toIndentedString(queues)).append("\n");
    sb.append("    agents: ").append(toIndentedString(agents)).append("\n");
    sb.append("    namespaces: ").append(toIndentedString(namespaces)).append("\n");
    sb.append("    userAccesses: ").append(toIndentedString(userAccesses)).append("\n");
    sb.append("    teams: ").append(toIndentedString(teams)).append("\n");
    sb.append("    projects: ").append(toIndentedString(projects)).append("\n");
    sb.append("    policy: ").append(toIndentedString(policy)).append("\n");
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
    openapiFields.add("connections");
    openapiFields.add("preset");
    openapiFields.add("presets");
    openapiFields.add("queue");
    openapiFields.add("queues");
    openapiFields.add("agents");
    openapiFields.add("namespaces");
    openapiFields.add("user_accesses");
    openapiFields.add("teams");
    openapiFields.add("projects");
    openapiFields.add("policy");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1ProjectSettings
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1ProjectSettings.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1ProjectSettings is not found in the empty JSON string", V1ProjectSettings.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1ProjectSettings.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1ProjectSettings` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("connections") != null && !jsonObj.get("connections").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `connections` to be an array in the JSON string but got `%s`", jsonObj.get("connections").toString()));
      }
      if ((jsonObj.get("preset") != null && !jsonObj.get("preset").isJsonNull()) && !jsonObj.get("preset").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `preset` to be a primitive type in the JSON string but got `%s`", jsonObj.get("preset").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("presets") != null && !jsonObj.get("presets").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `presets` to be an array in the JSON string but got `%s`", jsonObj.get("presets").toString()));
      }
      if ((jsonObj.get("queue") != null && !jsonObj.get("queue").isJsonNull()) && !jsonObj.get("queue").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `queue` to be a primitive type in the JSON string but got `%s`", jsonObj.get("queue").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("queues") != null && !jsonObj.get("queues").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `queues` to be an array in the JSON string but got `%s`", jsonObj.get("queues").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("agents") != null && !jsonObj.get("agents").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `agents` to be an array in the JSON string but got `%s`", jsonObj.get("agents").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("namespaces") != null && !jsonObj.get("namespaces").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `namespaces` to be an array in the JSON string but got `%s`", jsonObj.get("namespaces").toString()));
      }
      if (jsonObj.get("user_accesses") != null && !jsonObj.get("user_accesses").isJsonNull()) {
        JsonArray jsonArrayuserAccesses = jsonObj.getAsJsonArray("user_accesses");
        if (jsonArrayuserAccesses != null) {
          // ensure the json data is an array
          if (!jsonObj.get("user_accesses").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `user_accesses` to be an array in the JSON string but got `%s`", jsonObj.get("user_accesses").toString()));
          }

          // validate the optional field `user_accesses` (array)
          for (int i = 0; i < jsonArrayuserAccesses.size(); i++) {
            V1UserAccess.validateJsonObject(jsonArrayuserAccesses.get(i).getAsJsonObject());
          };
        }
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("teams") != null && !jsonObj.get("teams").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `teams` to be an array in the JSON string but got `%s`", jsonObj.get("teams").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("projects") != null && !jsonObj.get("projects").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `projects` to be an array in the JSON string but got `%s`", jsonObj.get("projects").toString()));
      }
      if ((jsonObj.get("policy") != null && !jsonObj.get("policy").isJsonNull()) && !jsonObj.get("policy").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `policy` to be a primitive type in the JSON string but got `%s`", jsonObj.get("policy").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1ProjectSettings.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1ProjectSettings' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1ProjectSettings> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1ProjectSettings.class));

       return (TypeAdapter<T>) new TypeAdapter<V1ProjectSettings>() {
           @Override
           public void write(JsonWriter out, V1ProjectSettings value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1ProjectSettings read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1ProjectSettings given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1ProjectSettings
  * @throws IOException if the JSON string is invalid with respect to V1ProjectSettings
  */
  public static V1ProjectSettings fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1ProjectSettings.class);
  }

 /**
  * Convert an instance of V1ProjectSettings to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

