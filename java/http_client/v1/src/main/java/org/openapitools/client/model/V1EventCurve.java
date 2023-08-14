/*
 * Polyaxon SDKs and REST API specification.
 *
 *
 * The version of the OpenAPI document: 2.0.0-rc36
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
import org.openapitools.client.model.V1EventCurveKind;

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
 * V1EventCurve
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1EventCurve {
  public static final String SERIALIZED_NAME_KIND = "kind";
  @SerializedName(SERIALIZED_NAME_KIND)
  private V1EventCurveKind kind = V1EventCurveKind.ROC;

  public static final String SERIALIZED_NAME_X = "x";
  @SerializedName(SERIALIZED_NAME_X)
  private List<Float> x;

  public static final String SERIALIZED_NAME_Y = "y";
  @SerializedName(SERIALIZED_NAME_Y)
  private List<Float> y;

  public static final String SERIALIZED_NAME_ANNOTATION = "annotation";
  @SerializedName(SERIALIZED_NAME_ANNOTATION)
  private String annotation;

  public V1EventCurve() {
  }

  public V1EventCurve kind(V1EventCurveKind kind) {

    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @javax.annotation.Nullable

  public V1EventCurveKind getKind() {
    return kind;
  }


  public void setKind(V1EventCurveKind kind) {
    this.kind = kind;
  }


  public V1EventCurve x(List<Float> x) {

    this.x = x;
    return this;
  }

  public V1EventCurve addXItem(Float xItem) {
    if (this.x == null) {
      this.x = new ArrayList<>();
    }
    this.x.add(xItem);
    return this;
  }

   /**
   * Get x
   * @return x
  **/
  @javax.annotation.Nullable

  public List<Float> getX() {
    return x;
  }


  public void setX(List<Float> x) {
    this.x = x;
  }


  public V1EventCurve y(List<Float> y) {

    this.y = y;
    return this;
  }

  public V1EventCurve addYItem(Float yItem) {
    if (this.y == null) {
      this.y = new ArrayList<>();
    }
    this.y.add(yItem);
    return this;
  }

   /**
   * Get y
   * @return y
  **/
  @javax.annotation.Nullable

  public List<Float> getY() {
    return y;
  }


  public void setY(List<Float> y) {
    this.y = y;
  }


  public V1EventCurve annotation(String annotation) {

    this.annotation = annotation;
    return this;
  }

   /**
   * Get annotation
   * @return annotation
  **/
  @javax.annotation.Nullable

  public String getAnnotation() {
    return annotation;
  }


  public void setAnnotation(String annotation) {
    this.annotation = annotation;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1EventCurve v1EventCurve = (V1EventCurve) o;
    return Objects.equals(this.kind, v1EventCurve.kind) &&
        Objects.equals(this.x, v1EventCurve.x) &&
        Objects.equals(this.y, v1EventCurve.y) &&
        Objects.equals(this.annotation, v1EventCurve.annotation);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, x, y, annotation);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1EventCurve {\n");
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    x: ").append(toIndentedString(x)).append("\n");
    sb.append("    y: ").append(toIndentedString(y)).append("\n");
    sb.append("    annotation: ").append(toIndentedString(annotation)).append("\n");
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
    openapiFields.add("kind");
    openapiFields.add("x");
    openapiFields.add("y");
    openapiFields.add("annotation");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1EventCurve
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1EventCurve.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1EventCurve is not found in the empty JSON string", V1EventCurve.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1EventCurve.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1EventCurve` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("x") != null && !jsonObj.get("x").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `x` to be an array in the JSON string but got `%s`", jsonObj.get("x").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("y") != null && !jsonObj.get("y").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `y` to be an array in the JSON string but got `%s`", jsonObj.get("y").toString()));
      }
      if ((jsonObj.get("annotation") != null && !jsonObj.get("annotation").isJsonNull()) && !jsonObj.get("annotation").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `annotation` to be a primitive type in the JSON string but got `%s`", jsonObj.get("annotation").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1EventCurve.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1EventCurve' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1EventCurve> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1EventCurve.class));

       return (TypeAdapter<T>) new TypeAdapter<V1EventCurve>() {
           @Override
           public void write(JsonWriter out, V1EventCurve value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1EventCurve read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1EventCurve given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1EventCurve
  * @throws IOException if the JSON string is invalid with respect to V1EventCurve
  */
  public static V1EventCurve fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1EventCurve.class);
  }

 /**
  * Convert an instance of V1EventCurve to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

