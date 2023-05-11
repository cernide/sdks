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
 * The version of the OpenAPI document: 2.0.0-rc14
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
import org.openapitools.client.model.AgentStateResponseAgentState;
import org.openapitools.client.model.V1Statuses;

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
 * V1AgentStateResponse
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1AgentStateResponse {
  public static final String SERIALIZED_NAME_STATUS = "status";
  @SerializedName(SERIALIZED_NAME_STATUS)
  private V1Statuses status = V1Statuses.CREATED;

  public static final String SERIALIZED_NAME_STATE = "state";
  @SerializedName(SERIALIZED_NAME_STATE)
  private AgentStateResponseAgentState state;

  public static final String SERIALIZED_NAME_LIVE_STATE = "live_state";
  @SerializedName(SERIALIZED_NAME_LIVE_STATE)
  private Integer liveState;

  public static final String SERIALIZED_NAME_COMPATIBLE_UPDATES = "compatible_updates";
  @SerializedName(SERIALIZED_NAME_COMPATIBLE_UPDATES)
  private Object compatibleUpdates;

  public V1AgentStateResponse() {
  }

  public V1AgentStateResponse status(V1Statuses status) {

    this.status = status;
    return this;
  }

   /**
   * Get status
   * @return status
  **/
  @javax.annotation.Nullable

  public V1Statuses getStatus() {
    return status;
  }


  public void setStatus(V1Statuses status) {
    this.status = status;
  }


  public V1AgentStateResponse state(AgentStateResponseAgentState state) {

    this.state = state;
    return this;
  }

   /**
   * Get state
   * @return state
  **/
  @javax.annotation.Nullable

  public AgentStateResponseAgentState getState() {
    return state;
  }


  public void setState(AgentStateResponseAgentState state) {
    this.state = state;
  }


  public V1AgentStateResponse liveState(Integer liveState) {

    this.liveState = liveState;
    return this;
  }

   /**
   * Get liveState
   * @return liveState
  **/
  @javax.annotation.Nullable

  public Integer getLiveState() {
    return liveState;
  }


  public void setLiveState(Integer liveState) {
    this.liveState = liveState;
  }


  public V1AgentStateResponse compatibleUpdates(Object compatibleUpdates) {

    this.compatibleUpdates = compatibleUpdates;
    return this;
  }

   /**
   * Get compatibleUpdates
   * @return compatibleUpdates
  **/
  @javax.annotation.Nullable

  public Object getCompatibleUpdates() {
    return compatibleUpdates;
  }


  public void setCompatibleUpdates(Object compatibleUpdates) {
    this.compatibleUpdates = compatibleUpdates;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1AgentStateResponse v1AgentStateResponse = (V1AgentStateResponse) o;
    return Objects.equals(this.status, v1AgentStateResponse.status) &&
        Objects.equals(this.state, v1AgentStateResponse.state) &&
        Objects.equals(this.liveState, v1AgentStateResponse.liveState) &&
        Objects.equals(this.compatibleUpdates, v1AgentStateResponse.compatibleUpdates);
  }

  @Override
  public int hashCode() {
    return Objects.hash(status, state, liveState, compatibleUpdates);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1AgentStateResponse {\n");
    sb.append("    status: ").append(toIndentedString(status)).append("\n");
    sb.append("    state: ").append(toIndentedString(state)).append("\n");
    sb.append("    liveState: ").append(toIndentedString(liveState)).append("\n");
    sb.append("    compatibleUpdates: ").append(toIndentedString(compatibleUpdates)).append("\n");
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
    openapiFields.add("status");
    openapiFields.add("state");
    openapiFields.add("live_state");
    openapiFields.add("compatible_updates");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1AgentStateResponse
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1AgentStateResponse.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1AgentStateResponse is not found in the empty JSON string", V1AgentStateResponse.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1AgentStateResponse.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1AgentStateResponse` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `state`
      if (jsonObj.get("state") != null && !jsonObj.get("state").isJsonNull()) {
        AgentStateResponseAgentState.validateJsonObject(jsonObj.getAsJsonObject("state"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1AgentStateResponse.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1AgentStateResponse' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1AgentStateResponse> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1AgentStateResponse.class));

       return (TypeAdapter<T>) new TypeAdapter<V1AgentStateResponse>() {
           @Override
           public void write(JsonWriter out, V1AgentStateResponse value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1AgentStateResponse read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1AgentStateResponse given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1AgentStateResponse
  * @throws IOException if the JSON string is invalid with respect to V1AgentStateResponse
  */
  public static V1AgentStateResponse fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1AgentStateResponse.class);
  }

 /**
  * Convert an instance of V1AgentStateResponse to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

