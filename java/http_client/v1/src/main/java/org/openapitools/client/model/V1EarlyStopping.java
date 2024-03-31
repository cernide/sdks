/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.1.4
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
import org.openapitools.client.model.V1DiffStoppingPolicy;
import org.openapitools.client.model.V1FailureEarlyStopping;
import org.openapitools.client.model.V1MedianStoppingPolicy;
import org.openapitools.client.model.V1MetricEarlyStopping;
import org.openapitools.client.model.V1TruncationStoppingPolicy;

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
 * V1EarlyStopping
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1EarlyStopping {
  public static final String SERIALIZED_NAME_MEDIAN = "median";
  @SerializedName(SERIALIZED_NAME_MEDIAN)
  private V1MedianStoppingPolicy median;

  public static final String SERIALIZED_NAME_DIFF = "diff";
  @SerializedName(SERIALIZED_NAME_DIFF)
  private V1DiffStoppingPolicy diff;

  public static final String SERIALIZED_NAME_TRUNCATION = "truncation";
  @SerializedName(SERIALIZED_NAME_TRUNCATION)
  private V1TruncationStoppingPolicy truncation;

  public static final String SERIALIZED_NAME_METRIC = "metric";
  @SerializedName(SERIALIZED_NAME_METRIC)
  private V1MetricEarlyStopping metric;

  public static final String SERIALIZED_NAME_FAILURE = "failure";
  @SerializedName(SERIALIZED_NAME_FAILURE)
  private V1FailureEarlyStopping failure;

  public V1EarlyStopping() {
  }

  public V1EarlyStopping median(V1MedianStoppingPolicy median) {

    this.median = median;
    return this;
  }

   /**
   * Get median
   * @return median
  **/
  @javax.annotation.Nullable

  public V1MedianStoppingPolicy getMedian() {
    return median;
  }


  public void setMedian(V1MedianStoppingPolicy median) {
    this.median = median;
  }


  public V1EarlyStopping diff(V1DiffStoppingPolicy diff) {

    this.diff = diff;
    return this;
  }

   /**
   * Get diff
   * @return diff
  **/
  @javax.annotation.Nullable

  public V1DiffStoppingPolicy getDiff() {
    return diff;
  }


  public void setDiff(V1DiffStoppingPolicy diff) {
    this.diff = diff;
  }


  public V1EarlyStopping truncation(V1TruncationStoppingPolicy truncation) {

    this.truncation = truncation;
    return this;
  }

   /**
   * Get truncation
   * @return truncation
  **/
  @javax.annotation.Nullable

  public V1TruncationStoppingPolicy getTruncation() {
    return truncation;
  }


  public void setTruncation(V1TruncationStoppingPolicy truncation) {
    this.truncation = truncation;
  }


  public V1EarlyStopping metric(V1MetricEarlyStopping metric) {

    this.metric = metric;
    return this;
  }

   /**
   * Get metric
   * @return metric
  **/
  @javax.annotation.Nullable

  public V1MetricEarlyStopping getMetric() {
    return metric;
  }


  public void setMetric(V1MetricEarlyStopping metric) {
    this.metric = metric;
  }


  public V1EarlyStopping failure(V1FailureEarlyStopping failure) {

    this.failure = failure;
    return this;
  }

   /**
   * Get failure
   * @return failure
  **/
  @javax.annotation.Nullable

  public V1FailureEarlyStopping getFailure() {
    return failure;
  }


  public void setFailure(V1FailureEarlyStopping failure) {
    this.failure = failure;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1EarlyStopping v1EarlyStopping = (V1EarlyStopping) o;
    return Objects.equals(this.median, v1EarlyStopping.median) &&
        Objects.equals(this.diff, v1EarlyStopping.diff) &&
        Objects.equals(this.truncation, v1EarlyStopping.truncation) &&
        Objects.equals(this.metric, v1EarlyStopping.metric) &&
        Objects.equals(this.failure, v1EarlyStopping.failure);
  }

  @Override
  public int hashCode() {
    return Objects.hash(median, diff, truncation, metric, failure);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1EarlyStopping {\n");
    sb.append("    median: ").append(toIndentedString(median)).append("\n");
    sb.append("    diff: ").append(toIndentedString(diff)).append("\n");
    sb.append("    truncation: ").append(toIndentedString(truncation)).append("\n");
    sb.append("    metric: ").append(toIndentedString(metric)).append("\n");
    sb.append("    failure: ").append(toIndentedString(failure)).append("\n");
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
    openapiFields.add("median");
    openapiFields.add("diff");
    openapiFields.add("truncation");
    openapiFields.add("metric");
    openapiFields.add("failure");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1EarlyStopping
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1EarlyStopping.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1EarlyStopping is not found in the empty JSON string", V1EarlyStopping.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1EarlyStopping.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1EarlyStopping` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // validate the optional field `median`
      if (jsonObj.get("median") != null && !jsonObj.get("median").isJsonNull()) {
        V1MedianStoppingPolicy.validateJsonObject(jsonObj.getAsJsonObject("median"));
      }
      // validate the optional field `diff`
      if (jsonObj.get("diff") != null && !jsonObj.get("diff").isJsonNull()) {
        V1DiffStoppingPolicy.validateJsonObject(jsonObj.getAsJsonObject("diff"));
      }
      // validate the optional field `truncation`
      if (jsonObj.get("truncation") != null && !jsonObj.get("truncation").isJsonNull()) {
        V1TruncationStoppingPolicy.validateJsonObject(jsonObj.getAsJsonObject("truncation"));
      }
      // validate the optional field `metric`
      if (jsonObj.get("metric") != null && !jsonObj.get("metric").isJsonNull()) {
        V1MetricEarlyStopping.validateJsonObject(jsonObj.getAsJsonObject("metric"));
      }
      // validate the optional field `failure`
      if (jsonObj.get("failure") != null && !jsonObj.get("failure").isJsonNull()) {
        V1FailureEarlyStopping.validateJsonObject(jsonObj.getAsJsonObject("failure"));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1EarlyStopping.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1EarlyStopping' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1EarlyStopping> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1EarlyStopping.class));

       return (TypeAdapter<T>) new TypeAdapter<V1EarlyStopping>() {
           @Override
           public void write(JsonWriter out, V1EarlyStopping value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1EarlyStopping read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1EarlyStopping given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1EarlyStopping
  * @throws IOException if the JSON string is invalid with respect to V1EarlyStopping
  */
  public static V1EarlyStopping fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1EarlyStopping.class);
  }

 /**
  * Convert an instance of V1EarlyStopping to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

