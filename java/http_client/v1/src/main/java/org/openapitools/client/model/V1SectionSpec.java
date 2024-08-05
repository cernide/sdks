/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.3.2
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
 * V1SectionSpec
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1SectionSpec {
  public static final String SERIALIZED_NAME_NAME = "name";
  @SerializedName(SERIALIZED_NAME_NAME)
  private String name;

  public static final String SERIALIZED_NAME_IS_MINIMIZED = "is_minimized";
  @SerializedName(SERIALIZED_NAME_IS_MINIMIZED)
  private Boolean isMinimized;

  public static final String SERIALIZED_NAME_IS_FROZEN = "is_frozen";
  @SerializedName(SERIALIZED_NAME_IS_FROZEN)
  private Boolean isFrozen;

  public static final String SERIALIZED_NAME_COLUMNS = "columns";
  @SerializedName(SERIALIZED_NAME_COLUMNS)
  private Integer columns;

  public static final String SERIALIZED_NAME_HEIGHT = "height";
  @SerializedName(SERIALIZED_NAME_HEIGHT)
  private Integer height;

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

  public static final String SERIALIZED_NAME_WIDGETS = "widgets";
  @SerializedName(SERIALIZED_NAME_WIDGETS)
  private List<Object> widgets;

  public static final String SERIALIZED_NAME_PAGE_INDEX = "page_index";
  @SerializedName(SERIALIZED_NAME_PAGE_INDEX)
  private Integer pageIndex;

  public static final String SERIALIZED_NAME_PAGE_SIZE = "page_size";
  @SerializedName(SERIALIZED_NAME_PAGE_SIZE)
  private Integer pageSize;

  public V1SectionSpec() {
  }

  public V1SectionSpec name(String name) {

    this.name = name;
    return this;
  }

   /**
   * Get name
   * @return name
  **/
  @javax.annotation.Nullable

  public String getName() {
    return name;
  }


  public void setName(String name) {
    this.name = name;
  }


  public V1SectionSpec isMinimized(Boolean isMinimized) {

    this.isMinimized = isMinimized;
    return this;
  }

   /**
   * Get isMinimized
   * @return isMinimized
  **/
  @javax.annotation.Nullable

  public Boolean getIsMinimized() {
    return isMinimized;
  }


  public void setIsMinimized(Boolean isMinimized) {
    this.isMinimized = isMinimized;
  }


  public V1SectionSpec isFrozen(Boolean isFrozen) {

    this.isFrozen = isFrozen;
    return this;
  }

   /**
   * Get isFrozen
   * @return isFrozen
  **/
  @javax.annotation.Nullable

  public Boolean getIsFrozen() {
    return isFrozen;
  }


  public void setIsFrozen(Boolean isFrozen) {
    this.isFrozen = isFrozen;
  }


  public V1SectionSpec columns(Integer columns) {

    this.columns = columns;
    return this;
  }

   /**
   * Get columns
   * @return columns
  **/
  @javax.annotation.Nullable

  public Integer getColumns() {
    return columns;
  }


  public void setColumns(Integer columns) {
    this.columns = columns;
  }


  public V1SectionSpec height(Integer height) {

    this.height = height;
    return this;
  }

   /**
   * Get height
   * @return height
  **/
  @javax.annotation.Nullable

  public Integer getHeight() {
    return height;
  }


  public void setHeight(Integer height) {
    this.height = height;
  }


  public V1SectionSpec xaxis(String xaxis) {

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


  public V1SectionSpec smoothing(Integer smoothing) {

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


  public V1SectionSpec ignoreOutliers(Boolean ignoreOutliers) {

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


  public V1SectionSpec sampleSize(Integer sampleSize) {

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


  public V1SectionSpec widgets(List<Object> widgets) {

    this.widgets = widgets;
    return this;
  }

  public V1SectionSpec addWidgetsItem(Object widgetsItem) {
    if (this.widgets == null) {
      this.widgets = new ArrayList<>();
    }
    this.widgets.add(widgetsItem);
    return this;
  }

   /**
   * Get widgets
   * @return widgets
  **/
  @javax.annotation.Nullable

  public List<Object> getWidgets() {
    return widgets;
  }


  public void setWidgets(List<Object> widgets) {
    this.widgets = widgets;
  }


  public V1SectionSpec pageIndex(Integer pageIndex) {

    this.pageIndex = pageIndex;
    return this;
  }

   /**
   * Get pageIndex
   * @return pageIndex
  **/
  @javax.annotation.Nullable

  public Integer getPageIndex() {
    return pageIndex;
  }


  public void setPageIndex(Integer pageIndex) {
    this.pageIndex = pageIndex;
  }


  public V1SectionSpec pageSize(Integer pageSize) {

    this.pageSize = pageSize;
    return this;
  }

   /**
   * Get pageSize
   * @return pageSize
  **/
  @javax.annotation.Nullable

  public Integer getPageSize() {
    return pageSize;
  }


  public void setPageSize(Integer pageSize) {
    this.pageSize = pageSize;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1SectionSpec v1SectionSpec = (V1SectionSpec) o;
    return Objects.equals(this.name, v1SectionSpec.name) &&
        Objects.equals(this.isMinimized, v1SectionSpec.isMinimized) &&
        Objects.equals(this.isFrozen, v1SectionSpec.isFrozen) &&
        Objects.equals(this.columns, v1SectionSpec.columns) &&
        Objects.equals(this.height, v1SectionSpec.height) &&
        Objects.equals(this.xaxis, v1SectionSpec.xaxis) &&
        Objects.equals(this.smoothing, v1SectionSpec.smoothing) &&
        Objects.equals(this.ignoreOutliers, v1SectionSpec.ignoreOutliers) &&
        Objects.equals(this.sampleSize, v1SectionSpec.sampleSize) &&
        Objects.equals(this.widgets, v1SectionSpec.widgets) &&
        Objects.equals(this.pageIndex, v1SectionSpec.pageIndex) &&
        Objects.equals(this.pageSize, v1SectionSpec.pageSize);
  }

  @Override
  public int hashCode() {
    return Objects.hash(name, isMinimized, isFrozen, columns, height, xaxis, smoothing, ignoreOutliers, sampleSize, widgets, pageIndex, pageSize);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1SectionSpec {\n");
    sb.append("    name: ").append(toIndentedString(name)).append("\n");
    sb.append("    isMinimized: ").append(toIndentedString(isMinimized)).append("\n");
    sb.append("    isFrozen: ").append(toIndentedString(isFrozen)).append("\n");
    sb.append("    columns: ").append(toIndentedString(columns)).append("\n");
    sb.append("    height: ").append(toIndentedString(height)).append("\n");
    sb.append("    xaxis: ").append(toIndentedString(xaxis)).append("\n");
    sb.append("    smoothing: ").append(toIndentedString(smoothing)).append("\n");
    sb.append("    ignoreOutliers: ").append(toIndentedString(ignoreOutliers)).append("\n");
    sb.append("    sampleSize: ").append(toIndentedString(sampleSize)).append("\n");
    sb.append("    widgets: ").append(toIndentedString(widgets)).append("\n");
    sb.append("    pageIndex: ").append(toIndentedString(pageIndex)).append("\n");
    sb.append("    pageSize: ").append(toIndentedString(pageSize)).append("\n");
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
    openapiFields.add("name");
    openapiFields.add("is_minimized");
    openapiFields.add("is_frozen");
    openapiFields.add("columns");
    openapiFields.add("height");
    openapiFields.add("xaxis");
    openapiFields.add("smoothing");
    openapiFields.add("ignore_outliers");
    openapiFields.add("sample_size");
    openapiFields.add("widgets");
    openapiFields.add("page_index");
    openapiFields.add("page_size");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1SectionSpec
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1SectionSpec.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1SectionSpec is not found in the empty JSON string", V1SectionSpec.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1SectionSpec.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1SectionSpec` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("name") != null && !jsonObj.get("name").isJsonNull()) && !jsonObj.get("name").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `name` to be a primitive type in the JSON string but got `%s`", jsonObj.get("name").toString()));
      }
      if ((jsonObj.get("xaxis") != null && !jsonObj.get("xaxis").isJsonNull()) && !jsonObj.get("xaxis").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `xaxis` to be a primitive type in the JSON string but got `%s`", jsonObj.get("xaxis").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("widgets") != null && !jsonObj.get("widgets").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `widgets` to be an array in the JSON string but got `%s`", jsonObj.get("widgets").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1SectionSpec.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1SectionSpec' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1SectionSpec> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1SectionSpec.class));

       return (TypeAdapter<T>) new TypeAdapter<V1SectionSpec>() {
           @Override
           public void write(JsonWriter out, V1SectionSpec value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1SectionSpec read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1SectionSpec given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1SectionSpec
  * @throws IOException if the JSON string is invalid with respect to V1SectionSpec
  */
  public static V1SectionSpec fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1SectionSpec.class);
  }

 /**
  * Convert an instance of V1SectionSpec to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

