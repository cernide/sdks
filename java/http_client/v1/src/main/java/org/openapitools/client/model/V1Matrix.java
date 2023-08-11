/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc33
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
import org.openapitools.client.model.V1Bayes;
import org.openapitools.client.model.V1GridSearch;
import org.openapitools.client.model.V1Hyperband;
import org.openapitools.client.model.V1Hyperopt;
import org.openapitools.client.model.V1Iterative;
import org.openapitools.client.model.V1Mapping;
import org.openapitools.client.model.V1RandomSearch;

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
 * V1Matrix
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Matrix {
  public static final String SERIALIZED_NAME_RANDOM = "random";
  @SerializedName(SERIALIZED_NAME_RANDOM)
  private V1RandomSearch random;

  public static final String SERIALIZED_NAME_GRID = "grid";
  @SerializedName(SERIALIZED_NAME_GRID)
  private V1GridSearch grid;

  public static final String SERIALIZED_NAME_HYPERBAND = "hyperband";
  @SerializedName(SERIALIZED_NAME_HYPERBAND)
  private V1Hyperband hyperband;

  public static final String SERIALIZED_NAME_BAYES = "bayes";
  @SerializedName(SERIALIZED_NAME_BAYES)
  private V1Bayes bayes;

  public static final String SERIALIZED_NAME_HYPEROPT = "hyperopt";
  @SerializedName(SERIALIZED_NAME_HYPEROPT)
  private V1Hyperopt hyperopt;

  public static final String SERIALIZED_NAME_ITERATIVE = "iterative";
  @SerializedName(SERIALIZED_NAME_ITERATIVE)
  private V1Iterative iterative;

  public static final String SERIALIZED_NAME_MAPPING = "mapping";
  @SerializedName(SERIALIZED_NAME_MAPPING)
  private V1Mapping mapping;

  public V1Matrix() {
  }

  public V1Matrix random(V1RandomSearch random) {

    this.random = random;
    return this;
  }

   /**
   * Get random
   * @return random
  **/
  @javax.annotation.Nullable

  public V1RandomSearch getRandom() {
    return random;
  }


  public void setRandom(V1RandomSearch random) {
    this.random = random;
  }


  public V1Matrix grid(V1GridSearch grid) {

    this.grid = grid;
    return this;
  }

   /**
   * Get grid
   * @return grid
  **/
  @javax.annotation.Nullable

  public V1GridSearch getGrid() {
    return grid;
  }


  public void setGrid(V1GridSearch grid) {
    this.grid = grid;
  }


  public V1Matrix hyperband(V1Hyperband hyperband) {

    this.hyperband = hyperband;
    return this;
  }

   /**
   * Get hyperband
   * @return hyperband
  **/
  @javax.annotation.Nullable

  public V1Hyperband getHyperband() {
    return hyperband;
  }


  public void setHyperband(V1Hyperband hyperband) {
    this.hyperband = hyperband;
  }


  public V1Matrix bayes(V1Bayes bayes) {

    this.bayes = bayes;
    return this;
  }

   /**
   * Get bayes
   * @return bayes
  **/
  @javax.annotation.Nullable

  public V1Bayes getBayes() {
    return bayes;
  }


  public void setBayes(V1Bayes bayes) {
    this.bayes = bayes;
  }


  public V1Matrix hyperopt(V1Hyperopt hyperopt) {

    this.hyperopt = hyperopt;
    return this;
  }

   /**
   * Get hyperopt
   * @return hyperopt
  **/
  @javax.annotation.Nullable

  public V1Hyperopt getHyperopt() {
    return hyperopt;
  }


  public void setHyperopt(V1Hyperopt hyperopt) {
    this.hyperopt = hyperopt;
  }


  public V1Matrix iterative(V1Iterative iterative) {

    this.iterative = iterative;
    return this;
  }

   /**
   * Get iterative
   * @return iterative
  **/
  @javax.annotation.Nullable

  public V1Iterative getIterative() {
    return iterative;
  }


  public void setIterative(V1Iterative iterative) {
    this.iterative = iterative;
  }


  public V1Matrix mapping(V1Mapping mapping) {

    this.mapping = mapping;
    return this;
  }

   /**
   * Get mapping
   * @return mapping
  **/
  @javax.annotation.Nullable

  public V1Mapping getMapping() {
    return mapping;
  }


  public void setMapping(V1Mapping mapping) {
    this.mapping = mapping;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Matrix v1Matrix = (V1Matrix) o;
    return Objects.equals(this.random, v1Matrix.random) &&
        Objects.equals(this.grid, v1Matrix.grid) &&
        Objects.equals(this.hyperband, v1Matrix.hyperband) &&
        Objects.equals(this.bayes, v1Matrix.bayes) &&
        Objects.equals(this.hyperopt, v1Matrix.hyperopt) &&
        Objects.equals(this.iterative, v1Matrix.iterative) &&
        Objects.equals(this.mapping, v1Matrix.mapping);
  }

  @Override
  public int hashCode() {
    return Objects.hash(random, grid, hyperband, bayes, hyperopt, iterative, mapping);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Matrix {\n");
    sb.append("    random: ").append(toIndentedString(random)).append("\n");
    sb.append("    grid: ").append(toIndentedString(grid)).append("\n");
    sb.append("    hyperband: ").append(toIndentedString(hyperband)).append("\n");
    sb.append("    bayes: ").append(toIndentedString(bayes)).append("\n");
    sb.append("    hyperopt: ").append(toIndentedString(hyperopt)).append("\n");
    sb.append("    iterative: ").append(toIndentedString(iterative)).append("\n");
    sb.append("    mapping: ").append(toIndentedString(mapping)).append("\n");
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
    openapiFields.add("random");
    openapiFields.add("grid");
    openapiFields.add("hyperband");
    openapiFields.add("bayes");
    openapiFields.add("hyperopt");
    openapiFields.add("iterative");
    openapiFields.add("mapping");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Matrix
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Matrix.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Matrix is not found in the empty JSON string", V1Matrix.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Matrix.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Matrix` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `random`
      if (jsonObj.get("random") != null && !jsonObj.get("random").isJsonNull()) {
        V1RandomSearch.validateJsonObject(jsonObj.getAsJsonObject("random"));
      }
      // validate the optional field `grid`
      if (jsonObj.get("grid") != null && !jsonObj.get("grid").isJsonNull()) {
        V1GridSearch.validateJsonObject(jsonObj.getAsJsonObject("grid"));
      }
      // validate the optional field `hyperband`
      if (jsonObj.get("hyperband") != null && !jsonObj.get("hyperband").isJsonNull()) {
        V1Hyperband.validateJsonObject(jsonObj.getAsJsonObject("hyperband"));
      }
      // validate the optional field `bayes`
      if (jsonObj.get("bayes") != null && !jsonObj.get("bayes").isJsonNull()) {
        V1Bayes.validateJsonObject(jsonObj.getAsJsonObject("bayes"));
      }
      // validate the optional field `hyperopt`
      if (jsonObj.get("hyperopt") != null && !jsonObj.get("hyperopt").isJsonNull()) {
        V1Hyperopt.validateJsonObject(jsonObj.getAsJsonObject("hyperopt"));
      }
      // validate the optional field `iterative`
      if (jsonObj.get("iterative") != null && !jsonObj.get("iterative").isJsonNull()) {
        V1Iterative.validateJsonObject(jsonObj.getAsJsonObject("iterative"));
      }
      // validate the optional field `mapping`
      if (jsonObj.get("mapping") != null && !jsonObj.get("mapping").isJsonNull()) {
        V1Mapping.validateJsonObject(jsonObj.getAsJsonObject("mapping"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Matrix.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Matrix' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Matrix> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Matrix.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Matrix>() {
           @Override
           public void write(JsonWriter out, V1Matrix value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Matrix read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Matrix given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Matrix
  * @throws IOException if the JSON string is invalid with respect to V1Matrix
  */
  public static V1Matrix fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Matrix.class);
  }

 /**
  * Convert an instance of V1Matrix to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

