// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'list.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

ImageList _$ImageListFromJson(Map<String, dynamic> json) => ImageList(
      token: json['token'] as String?,
      images:
          (json['images'] as List<dynamic>?)?.map((e) => e as String).toList(),
    );

Map<String, dynamic> _$ImageListToJson(ImageList instance) => <String, dynamic>{
      'token': instance.token,
      'images': instance.images,
    };
