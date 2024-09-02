/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.4.0
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
import org.openapitools.client.model.V1SectionSpec;

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
 * V1DashboardSpec
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1DashboardSpec {
  public static final String SERIALIZED_NAME_SECTIONS = "sections";
  @SerializedName(SERIALIZED_NAME_SECTIONS)
  private List<V1SectionSpec> sections;

  public static final String SERIALIZED_NAME_XAXIS = "xaxis";
  @SerializedName(SERIALIZED_NAME_XAXIS)
  private String xaxis;

  public static final String SERIALIZED_NAME_SMOOTHING = "smoothing";
  @SerializedName(SERIALIZED_NAME_SMOOTHING)
  private Integer smoothing;

  public static final String SERIALIZED_NAME_IGNORE_OUTLIERS = "ignore_outliers";
  @SerializedName(SERIALIZED_NAME_IGNORE_OUTLIERS)
  private Boolean ignoreOutliers;

  public static final String SERIALIZED_NAME_SAMPLE_SIZE = "sample_size";
  @SerializedName(SERIALIZED_NAME_SAMPLE_SIZE)
  private Integer sampleSize;

  public V1DashboardSpec() {
  }

  public V1DashboardSpec sections(List<V1SectionSpec> sections) {

    this.sections = sections;
    return this;
  }

  public V1DashboardSpec addSectionsItem(V1SectionSpec sectionsItem) {
    if (this.sections == null) {
      this.sections = new ArrayList<>();
    }
    this.sections.add(sectionsItem);
    return this;
  }

   /**
   * Get sections
   * @return sections
  **/
  @javax.annotation.Nullable

  public List<V1SectionSpec> getSections() {
    return sections;
  }


  public void setSections(List<V1SectionSpec> sections) {
    this.sections = sections;
  }


  public V1DashboardSpec xaxis(String xaxis) {

    this.xaxis = xaxis;
    return this;
  }

   /**
   * Get xaxis
   * @return xaxis
  **/
  @javax.annotation.Nullable

  public String getXaxis() {
    return xaxis;
  }


  public void setXaxis(String xaxis) {
    this.xaxis = xaxis;
  }


  public V1DashboardSpec smoothing(Integer smoothing) {

    this.smoothing = smoothing;
    return this;
  }

   /**
   * Get smoothing
   * @return smoothing
  **/
  @javax.annotation.Nullable

  public Integer getSmoothing() {
    return smoothing;
  }


  public void setSmoothing(Integer smoothing) {
    this.smoothing = smoothing;
  }


  public V1DashboardSpec ignoreOutliers(Boolean ignoreOutliers) {

    this.ignoreOutliers = ignoreOutliers;
    return this;
  }

   /**
   * Get ignoreOutliers
   * @return ignoreOutliers
  **/
  @javax.annotation.Nullable

  public Boolean getIgnoreOutliers() {
    return ignoreOutliers;
  }


  public void setIgnoreOutliers(Boolean ignoreOutliers) {
    this.ignoreOutliers = ignoreOutliers;
  }


  public V1DashboardSpec sampleSize(Integer sampleSize) {

    this.sampleSize = sampleSize;
    return this;
  }

   /**
   * Get sampleSize
   * @return sampleSize
  **/
  @javax.annotation.Nullable

  public Integer getSampleSize() {
    return sampleSize;
  }


  public void setSampleSize(Integer sampleSize) {
    this.sampleSize = sampleSize;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1DashboardSpec v1DashboardSpec = (V1DashboardSpec) o;
    return Objects.equals(this.sections, v1DashboardSpec.sections) &&
        Objects.equals(this.xaxis, v1DashboardSpec.xaxis) &&
        Objects.equals(this.smoothing, v1DashboardSpec.smoothing) &&
        Objects.equals(this.ignoreOutliers, v1DashboardSpec.ignoreOutliers) &&
        Objects.equals(this.sampleSize, v1DashboardSpec.sampleSize);
  }

  @Override
  public int hashCode() {
    return Objects.hash(sections, xaxis, smoothing, ignoreOutliers, sampleSize);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1DashboardSpec {\n");
    sb.append("    sections: ").append(toIndentedString(sections)).append("\n");
    sb.append("    xaxis: ").append(toIndentedString(xaxis)).append("\n");
    sb.append("    smoothing: ").append(toIndentedString(smoothing)).append("\n");
    sb.append("    ignoreOutliers: ").append(toIndentedString(ignoreOutliers)).append("\n");
    sb.append("    sampleSize: ").append(toIndentedString(sampleSize)).append("\n");
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
    openapiFields.add("sections");
    openapiFields.add("xaxis");
    openapiFields.add("smoothing");
    openapiFields.add("ignore_outliers");
    openapiFields.add("sample_size");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1DashboardSpec
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1DashboardSpec.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1DashboardSpec is not found in the empty JSON string", V1DashboardSpec.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1DashboardSpec.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1DashboardSpec` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if (jsonObj.get("sections") != null && !jsonObj.get("sections").isJsonNull()) {
        JsonArray jsonArraysections = jsonObj.getAsJsonArray("sections");
        if (jsonArraysections != null) {
          // ensure the json data is an array
          if (!jsonObj.get("sections").isJsonArray()) {
            throw new IllegalArgumentException(String.format("Expected the field `sections` to be an array in the JSON string but got `%s`", jsonObj.get("sections").toString()));
          }

          // validate the optional field `sections` (array)
          for (int i = 0; i < jsonArraysections.size(); i++) {
            V1SectionSpec.validateJsonObject(jsonArraysections.get(i).getAsJsonObject());
          };
        }
      }
      if ((jsonObj.get("xaxis") != null && !jsonObj.get("xaxis").isJsonNull()) && !jsonObj.get("xaxis").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `xaxis` to be a primitive type in the JSON string but got `%s`", jsonObj.get("xaxis").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1DashboardSpec.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1DashboardSpec' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1DashboardSpec> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1DashboardSpec.class));

       return (TypeAdapter<T>) new TypeAdapter<V1DashboardSpec>() {
           @Override
           public void write(JsonWriter out, V1DashboardSpec value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1DashboardSpec read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1DashboardSpec given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1DashboardSpec
  * @throws IOException if the JSON string is invalid with respect to V1DashboardSpec
  */
  public static V1DashboardSpec fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1DashboardSpec.class);
  }

 /**
  * Convert an instance of V1DashboardSpec to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

