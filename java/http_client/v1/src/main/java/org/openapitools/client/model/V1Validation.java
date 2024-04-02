/*
 * Polyaxon SDKs and REST API specification.
 *    
 *
 * The version of the OpenAPI document: 2.1.5
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
 * V1Validation
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen")
public class V1Validation {
  public static final String SERIALIZED_NAME_DELAY = "delay";
  @SerializedName(SERIALIZED_NAME_DELAY)
  private Boolean delay;

  public static final String SERIALIZED_NAME_GT = "gt";
  @SerializedName(SERIALIZED_NAME_GT)
  private Integer gt;

  public static final String SERIALIZED_NAME_GE = "ge";
  @SerializedName(SERIALIZED_NAME_GE)
  private Integer ge;

  public static final String SERIALIZED_NAME_LT = "lt";
  @SerializedName(SERIALIZED_NAME_LT)
  private Integer lt;

  public static final String SERIALIZED_NAME_LE = "le";
  @SerializedName(SERIALIZED_NAME_LE)
  private Integer le;

  public static final String SERIALIZED_NAME_MULTIPLE_OF = "multipleOf";
  @SerializedName(SERIALIZED_NAME_MULTIPLE_OF)
  private Integer multipleOf;

  public static final String SERIALIZED_NAME_MIN_DIGITS = "minDigits";
  @SerializedName(SERIALIZED_NAME_MIN_DIGITS)
  private Integer minDigits;

  public static final String SERIALIZED_NAME_MAX_DIGITS = "maxDigits";
  @SerializedName(SERIALIZED_NAME_MAX_DIGITS)
  private Integer maxDigits;

  public static final String SERIALIZED_NAME_DECIMAL_PLACES = "decimalPlaces";
  @SerializedName(SERIALIZED_NAME_DECIMAL_PLACES)
  private Integer decimalPlaces;

  public static final String SERIALIZED_NAME_REGEX = "regex";
  @SerializedName(SERIALIZED_NAME_REGEX)
  private String regex;

  public static final String SERIALIZED_NAME_MIN_LENGTH = "minLength";
  @SerializedName(SERIALIZED_NAME_MIN_LENGTH)
  private Integer minLength;

  public static final String SERIALIZED_NAME_MAX_LENGTH = "maxLength";
  @SerializedName(SERIALIZED_NAME_MAX_LENGTH)
  private Integer maxLength;

  public static final String SERIALIZED_NAME_CONTAINS = "contains";
  @SerializedName(SERIALIZED_NAME_CONTAINS)
  private List<Object> contains;

  public static final String SERIALIZED_NAME_EXCLUDES = "excludes";
  @SerializedName(SERIALIZED_NAME_EXCLUDES)
  private List<Object> excludes;

  public static final String SERIALIZED_NAME_OPTIONS = "options";
  @SerializedName(SERIALIZED_NAME_OPTIONS)
  private List<Object> options;

  public static final String SERIALIZED_NAME_MIN_ITEMS = "minItems";
  @SerializedName(SERIALIZED_NAME_MIN_ITEMS)
  private Integer minItems;

  public static final String SERIALIZED_NAME_MAX_ITEMS = "maxItems";
  @SerializedName(SERIALIZED_NAME_MAX_ITEMS)
  private Integer maxItems;

  public static final String SERIALIZED_NAME_KEYS = "keys";
  @SerializedName(SERIALIZED_NAME_KEYS)
  private List<String> keys;

  public static final String SERIALIZED_NAME_CONTAINS_KEYS = "containsKeys";
  @SerializedName(SERIALIZED_NAME_CONTAINS_KEYS)
  private List<String> containsKeys;

  public static final String SERIALIZED_NAME_EXCLUDES_KEYS = "excludesKeys";
  @SerializedName(SERIALIZED_NAME_EXCLUDES_KEYS)
  private List<String> excludesKeys;

  public V1Validation() {
  }

  public V1Validation delay(Boolean delay) {
    
    this.delay = delay;
    return this;
  }

   /**
   * Get delay
   * @return delay
  **/
  @javax.annotation.Nullable

  public Boolean getDelay() {
    return delay;
  }


  public void setDelay(Boolean delay) {
    this.delay = delay;
  }


  public V1Validation gt(Integer gt) {
    
    this.gt = gt;
    return this;
  }

   /**
   * Get gt
   * @return gt
  **/
  @javax.annotation.Nullable

  public Integer getGt() {
    return gt;
  }


  public void setGt(Integer gt) {
    this.gt = gt;
  }


  public V1Validation ge(Integer ge) {
    
    this.ge = ge;
    return this;
  }

   /**
   * Get ge
   * @return ge
  **/
  @javax.annotation.Nullable

  public Integer getGe() {
    return ge;
  }


  public void setGe(Integer ge) {
    this.ge = ge;
  }


  public V1Validation lt(Integer lt) {
    
    this.lt = lt;
    return this;
  }

   /**
   * Get lt
   * @return lt
  **/
  @javax.annotation.Nullable

  public Integer getLt() {
    return lt;
  }


  public void setLt(Integer lt) {
    this.lt = lt;
  }


  public V1Validation le(Integer le) {
    
    this.le = le;
    return this;
  }

   /**
   * Get le
   * @return le
  **/
  @javax.annotation.Nullable

  public Integer getLe() {
    return le;
  }


  public void setLe(Integer le) {
    this.le = le;
  }


  public V1Validation multipleOf(Integer multipleOf) {
    
    this.multipleOf = multipleOf;
    return this;
  }

   /**
   * Get multipleOf
   * @return multipleOf
  **/
  @javax.annotation.Nullable

  public Integer getMultipleOf() {
    return multipleOf;
  }


  public void setMultipleOf(Integer multipleOf) {
    this.multipleOf = multipleOf;
  }


  public V1Validation minDigits(Integer minDigits) {
    
    this.minDigits = minDigits;
    return this;
  }

   /**
   * Get minDigits
   * @return minDigits
  **/
  @javax.annotation.Nullable

  public Integer getMinDigits() {
    return minDigits;
  }


  public void setMinDigits(Integer minDigits) {
    this.minDigits = minDigits;
  }


  public V1Validation maxDigits(Integer maxDigits) {
    
    this.maxDigits = maxDigits;
    return this;
  }

   /**
   * Get maxDigits
   * @return maxDigits
  **/
  @javax.annotation.Nullable

  public Integer getMaxDigits() {
    return maxDigits;
  }


  public void setMaxDigits(Integer maxDigits) {
    this.maxDigits = maxDigits;
  }


  public V1Validation decimalPlaces(Integer decimalPlaces) {
    
    this.decimalPlaces = decimalPlaces;
    return this;
  }

   /**
   * Get decimalPlaces
   * @return decimalPlaces
  **/
  @javax.annotation.Nullable

  public Integer getDecimalPlaces() {
    return decimalPlaces;
  }


  public void setDecimalPlaces(Integer decimalPlaces) {
    this.decimalPlaces = decimalPlaces;
  }


  public V1Validation regex(String regex) {
    
    this.regex = regex;
    return this;
  }

   /**
   * Get regex
   * @return regex
  **/
  @javax.annotation.Nullable

  public String getRegex() {
    return regex;
  }


  public void setRegex(String regex) {
    this.regex = regex;
  }


  public V1Validation minLength(Integer minLength) {
    
    this.minLength = minLength;
    return this;
  }

   /**
   * Get minLength
   * @return minLength
  **/
  @javax.annotation.Nullable

  public Integer getMinLength() {
    return minLength;
  }


  public void setMinLength(Integer minLength) {
    this.minLength = minLength;
  }


  public V1Validation maxLength(Integer maxLength) {
    
    this.maxLength = maxLength;
    return this;
  }

   /**
   * Get maxLength
   * @return maxLength
  **/
  @javax.annotation.Nullable

  public Integer getMaxLength() {
    return maxLength;
  }


  public void setMaxLength(Integer maxLength) {
    this.maxLength = maxLength;
  }


  public V1Validation contains(List<Object> contains) {
    
    this.contains = contains;
    return this;
  }

  public V1Validation addContainsItem(Object containsItem) {
    if (this.contains == null) {
      this.contains = new ArrayList<>();
    }
    this.contains.add(containsItem);
    return this;
  }

   /**
   * Get contains
   * @return contains
  **/
  @javax.annotation.Nullable

  public List<Object> getContains() {
    return contains;
  }


  public void setContains(List<Object> contains) {
    this.contains = contains;
  }


  public V1Validation excludes(List<Object> excludes) {
    
    this.excludes = excludes;
    return this;
  }

  public V1Validation addExcludesItem(Object excludesItem) {
    if (this.excludes == null) {
      this.excludes = new ArrayList<>();
    }
    this.excludes.add(excludesItem);
    return this;
  }

   /**
   * Get excludes
   * @return excludes
  **/
  @javax.annotation.Nullable

  public List<Object> getExcludes() {
    return excludes;
  }


  public void setExcludes(List<Object> excludes) {
    this.excludes = excludes;
  }


  public V1Validation options(List<Object> options) {
    
    this.options = options;
    return this;
  }

  public V1Validation addOptionsItem(Object optionsItem) {
    if (this.options == null) {
      this.options = new ArrayList<>();
    }
    this.options.add(optionsItem);
    return this;
  }

   /**
   * Get options
   * @return options
  **/
  @javax.annotation.Nullable

  public List<Object> getOptions() {
    return options;
  }


  public void setOptions(List<Object> options) {
    this.options = options;
  }


  public V1Validation minItems(Integer minItems) {
    
    this.minItems = minItems;
    return this;
  }

   /**
   * Get minItems
   * @return minItems
  **/
  @javax.annotation.Nullable

  public Integer getMinItems() {
    return minItems;
  }


  public void setMinItems(Integer minItems) {
    this.minItems = minItems;
  }


  public V1Validation maxItems(Integer maxItems) {
    
    this.maxItems = maxItems;
    return this;
  }

   /**
   * Get maxItems
   * @return maxItems
  **/
  @javax.annotation.Nullable

  public Integer getMaxItems() {
    return maxItems;
  }


  public void setMaxItems(Integer maxItems) {
    this.maxItems = maxItems;
  }


  public V1Validation keys(List<String> keys) {
    
    this.keys = keys;
    return this;
  }

  public V1Validation addKeysItem(String keysItem) {
    if (this.keys == null) {
      this.keys = new ArrayList<>();
    }
    this.keys.add(keysItem);
    return this;
  }

   /**
   * Get keys
   * @return keys
  **/
  @javax.annotation.Nullable

  public List<String> getKeys() {
    return keys;
  }


  public void setKeys(List<String> keys) {
    this.keys = keys;
  }


  public V1Validation containsKeys(List<String> containsKeys) {
    
    this.containsKeys = containsKeys;
    return this;
  }

  public V1Validation addContainsKeysItem(String containsKeysItem) {
    if (this.containsKeys == null) {
      this.containsKeys = new ArrayList<>();
    }
    this.containsKeys.add(containsKeysItem);
    return this;
  }

   /**
   * Get containsKeys
   * @return containsKeys
  **/
  @javax.annotation.Nullable

  public List<String> getContainsKeys() {
    return containsKeys;
  }


  public void setContainsKeys(List<String> containsKeys) {
    this.containsKeys = containsKeys;
  }


  public V1Validation excludesKeys(List<String> excludesKeys) {
    
    this.excludesKeys = excludesKeys;
    return this;
  }

  public V1Validation addExcludesKeysItem(String excludesKeysItem) {
    if (this.excludesKeys == null) {
      this.excludesKeys = new ArrayList<>();
    }
    this.excludesKeys.add(excludesKeysItem);
    return this;
  }

   /**
   * Get excludesKeys
   * @return excludesKeys
  **/
  @javax.annotation.Nullable

  public List<String> getExcludesKeys() {
    return excludesKeys;
  }


  public void setExcludesKeys(List<String> excludesKeys) {
    this.excludesKeys = excludesKeys;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1Validation v1Validation = (V1Validation) o;
    return Objects.equals(this.delay, v1Validation.delay) &&
        Objects.equals(this.gt, v1Validation.gt) &&
        Objects.equals(this.ge, v1Validation.ge) &&
        Objects.equals(this.lt, v1Validation.lt) &&
        Objects.equals(this.le, v1Validation.le) &&
        Objects.equals(this.multipleOf, v1Validation.multipleOf) &&
        Objects.equals(this.minDigits, v1Validation.minDigits) &&
        Objects.equals(this.maxDigits, v1Validation.maxDigits) &&
        Objects.equals(this.decimalPlaces, v1Validation.decimalPlaces) &&
        Objects.equals(this.regex, v1Validation.regex) &&
        Objects.equals(this.minLength, v1Validation.minLength) &&
        Objects.equals(this.maxLength, v1Validation.maxLength) &&
        Objects.equals(this.contains, v1Validation.contains) &&
        Objects.equals(this.excludes, v1Validation.excludes) &&
        Objects.equals(this.options, v1Validation.options) &&
        Objects.equals(this.minItems, v1Validation.minItems) &&
        Objects.equals(this.maxItems, v1Validation.maxItems) &&
        Objects.equals(this.keys, v1Validation.keys) &&
        Objects.equals(this.containsKeys, v1Validation.containsKeys) &&
        Objects.equals(this.excludesKeys, v1Validation.excludesKeys);
  }

  @Override
  public int hashCode() {
    return Objects.hash(delay, gt, ge, lt, le, multipleOf, minDigits, maxDigits, decimalPlaces, regex, minLength, maxLength, contains, excludes, options, minItems, maxItems, keys, containsKeys, excludesKeys);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1Validation {\n");
    sb.append("    delay: ").append(toIndentedString(delay)).append("\n");
    sb.append("    gt: ").append(toIndentedString(gt)).append("\n");
    sb.append("    ge: ").append(toIndentedString(ge)).append("\n");
    sb.append("    lt: ").append(toIndentedString(lt)).append("\n");
    sb.append("    le: ").append(toIndentedString(le)).append("\n");
    sb.append("    multipleOf: ").append(toIndentedString(multipleOf)).append("\n");
    sb.append("    minDigits: ").append(toIndentedString(minDigits)).append("\n");
    sb.append("    maxDigits: ").append(toIndentedString(maxDigits)).append("\n");
    sb.append("    decimalPlaces: ").append(toIndentedString(decimalPlaces)).append("\n");
    sb.append("    regex: ").append(toIndentedString(regex)).append("\n");
    sb.append("    minLength: ").append(toIndentedString(minLength)).append("\n");
    sb.append("    maxLength: ").append(toIndentedString(maxLength)).append("\n");
    sb.append("    contains: ").append(toIndentedString(contains)).append("\n");
    sb.append("    excludes: ").append(toIndentedString(excludes)).append("\n");
    sb.append("    options: ").append(toIndentedString(options)).append("\n");
    sb.append("    minItems: ").append(toIndentedString(minItems)).append("\n");
    sb.append("    maxItems: ").append(toIndentedString(maxItems)).append("\n");
    sb.append("    keys: ").append(toIndentedString(keys)).append("\n");
    sb.append("    containsKeys: ").append(toIndentedString(containsKeys)).append("\n");
    sb.append("    excludesKeys: ").append(toIndentedString(excludesKeys)).append("\n");
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
    openapiFields.add("delay");
    openapiFields.add("gt");
    openapiFields.add("ge");
    openapiFields.add("lt");
    openapiFields.add("le");
    openapiFields.add("multipleOf");
    openapiFields.add("minDigits");
    openapiFields.add("maxDigits");
    openapiFields.add("decimalPlaces");
    openapiFields.add("regex");
    openapiFields.add("minLength");
    openapiFields.add("maxLength");
    openapiFields.add("contains");
    openapiFields.add("excludes");
    openapiFields.add("options");
    openapiFields.add("minItems");
    openapiFields.add("maxItems");
    openapiFields.add("keys");
    openapiFields.add("containsKeys");
    openapiFields.add("excludesKeys");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to V1Validation
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!V1Validation.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in V1Validation is not found in the empty JSON string", V1Validation.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!V1Validation.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `V1Validation` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("regex") != null && !jsonObj.get("regex").isJsonNull()) && !jsonObj.get("regex").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `regex` to be a primitive type in the JSON string but got `%s`", jsonObj.get("regex").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("contains") != null && !jsonObj.get("contains").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `contains` to be an array in the JSON string but got `%s`", jsonObj.get("contains").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("excludes") != null && !jsonObj.get("excludes").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `excludes` to be an array in the JSON string but got `%s`", jsonObj.get("excludes").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("options") != null && !jsonObj.get("options").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `options` to be an array in the JSON string but got `%s`", jsonObj.get("options").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("keys") != null && !jsonObj.get("keys").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `keys` to be an array in the JSON string but got `%s`", jsonObj.get("keys").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("containsKeys") != null && !jsonObj.get("containsKeys").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `containsKeys` to be an array in the JSON string but got `%s`", jsonObj.get("containsKeys").toString()));
      }
      // ensure the optional json data is an array if present
      if (jsonObj.get("excludesKeys") != null && !jsonObj.get("excludesKeys").isJsonArray()) {
        throw new IllegalArgumentException(String.format("Expected the field `excludesKeys` to be an array in the JSON string but got `%s`", jsonObj.get("excludesKeys").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!V1Validation.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'V1Validation' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<V1Validation> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(V1Validation.class));

       return (TypeAdapter<T>) new TypeAdapter<V1Validation>() {
           @Override
           public void write(JsonWriter out, V1Validation value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public V1Validation read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of V1Validation given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of V1Validation
  * @throws IOException if the JSON string is invalid with respect to V1Validation
  */
  public static V1Validation fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, V1Validation.class);
  }

 /**
  * Convert an instance of V1Validation to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

