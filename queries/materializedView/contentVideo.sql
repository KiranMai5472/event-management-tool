CREATE VIEW content_Video AS

SELECT

    c.id AS "id",

    c.title AS "title",

    c.tagline AS "tagline",

    c.description AS "description",

    c.overview AS "overview",

    c.duration AS "duration",

    c.runtime AS "runtime",

    c.is_free AS "isFree",

    ct.content_type_name AS "contentType",

    cst.content_subtype_name AS "subContentType",

    pt.name AS "publishedType",

    TO_CHAR(c.published_date, 'DD-MM-YYYY') AS "publishedDate",

    l.name AS "language",
	
	jsonb_build_object(

        'name', pr.parental_rating_name,

        'shortName', pr.parental_rating_short_name,

        'tagline', pr.parental_rating_tagline,

        'description', pr.parental_rating_description,

        'warningTitle', pr.parental_rating_warning_title,

        'isKid', pr.is_kid,

        'image', jsonb_build_object(

            'type', it.name,

            'profile', pi.profile,

            'imageLink', pi.image_link,

            'layoutType', pi.layout_type

        )

    ) AS "parentalRating",
	
	jsonb_build_object(

        'name', dl.deeplink_name,

        'link', dl.deeplink_link

    ) AS "deeplink",
	
	 jsonb_build_object(

        'parentContentId', sm.parent_content_id,

        'subParentContentId', sm.sub_parent_content_id,

        'episodeNumber', sm.episode_number,

        'index', sm.index,

        'popularity', sm.popularity

    ) AS "seriesMeta",

    jsonb_agg(
    jsonb_build_object(
        'videoURL', jsonb_build_object(
            'type', vut.video_url_type_name,
            'profile', vup.video_url_profile_name,
            'drm_enabled', vu.drm_enabled,
            'licenceUrl', vu.licence_url,
            'contentUrl', vu.content_url,
            'protocol', vupr.video_url_protocol_name,
            'encryptionType', vue.video_url_encryption_type_name,
            'qualityBitrate', (
                SELECT jsonb_agg(
                    jsonb_build_object(
                        'name', qb.quality_bitrate_name,
                        'displayName', qb.quality_bitrate_display_name,
                        'minRange', qb.min_range,
                        'maxRange', qb.max_range
                    )
                )
                FROM quality_bitrate qb
                WHERE qb.video_url_id = vu.id 
            )
        )
    )
) AS "videoURL",
	
 jsonb_agg(
	 jsonb_build_object(
		'language', st.language,
		'mimeTypeId', smt.name,
		'url', st.url	 
	)
	) AS "subtitle",
	
 jsonb_build_object(

        'adUrl', adi.ad_url,

        'provider', adi.ad_provider,

        'cuePoints', adi.cue

    ) AS "adInfo",
	
jsonb_build_object(

        'videoId', ev.video_id,

        'key', ev.key,

        'player', jsonb_build_object(

            'playerIdentifier', ep.player_identifier,

            'playerKey', ep.player_key,

            'playerName', ep.player_name

        )

    ) AS "externalVideo",
	
jsonb_build_object(

        'displayName', ctry.display_name,
        'key', ctry.name,
		'imageLink',ctry.image_link,
		'shortName',ctry.short_name,
		'phoneCode',ctry.phone_code,
		'timezone',ctry.timezone
	
	)AS "country"
	
FROM

    content c

LEFT JOIN

    content_type ct ON c.content_type_id = ct.id

LEFT JOIN

    content_subtype cst ON c.content_subtype_id = cst.id

LEFT JOIN

    published_type pt ON c.published_type_id = pt.id

LEFT JOIN

    language l ON c.language_id = l.id
	
LEFT JOIN

    deeplink dl ON c.id = dl.content_id
	
LEFT JOIN

    parental_rating pr ON c.parental_rating_id = pr.id

LEFT JOIN

    image pi ON pr.image_id = pi.id
	
LEFT JOIN 

	image_type it ON pi.type_id = it.id
	
LEFT JOIN 

	image ci ON ci.content_id = c.id
	
LEFT JOIN 

	image_type cit ON ci.type_id = cit.id
	
LEFT JOIN

    series_meta sm ON c.series_meta_id = sm.id
	
LEFT JOIN 

	video_url_content vuc ON c.id = vuc.content_id
LEFT JOIN 

 video_url vu ON vuc.video_url_id = vu.id

LEFT JOIN 

video_url_type vut ON vu.type_id = vut.id

LEFT JOIN

    video_url_protocol vupr ON vu.protocol_id = vupr.id

LEFT JOIN

    video_url_encryption_type vue ON vu.encryption_type_id = vue.id

LEFT JOIN

video_url_profile vup ON vu.profile_id = vup.id

LEFT JOIN

    quality_bitrate qb ON vu.id = qb.video_url_id
	
LEFT JOIN

	subtitle_content sc ON c.id = sc.content_id

LEFT JOIN 

 subtitle st ON sc.subtitle_id = st.id
 
 LEFT JOIN 
 
 subtitle_mime_type smt ON st.mime_type_id = smt.id

LEFT JOIN 

ad_info adi ON vu.ad_info_id = adi.id

LEFT JOIN 

external_video ev ON c.id = ev.content_id

LEFT JOIN

external_player ep ON ev.player_id = ep.id

LEFT JOIN 

country_content cc ON c.id = cc.content_id

LEFT JOIN 

country ctry ON cc.country_id = ctry.id

GROUP BY

    c.id,
	ct.content_type_name,
	cst.content_subtype_name,
	pt.name,
	l.name,
	pr.parental_rating_name,
	pr.parental_rating_short_name,
	pr.parental_rating_tagline,
	pr.parental_rating_description,
	pr.parental_rating_warning_title,
	pr.is_kid,
	it.name,
	pi.profile,
   pi.image_link,
   pi.layout_type,
   dl.deeplink_name,
dl.deeplink_link,
 sm.parent_content_id,
sm.sub_parent_content_id,
sm.episode_number,
sm.index,
sm.popularity,
adi.ad_url,
adi.ad_provider,
adi.cue,
ev.video_id,
ev.key,
ep.player_identifier,
ep.player_key,
ep.player_name,
ctry.display_name,
ctry.name,
ctry.image_link,
ctry.short_name,
ctry.phone_code,
ctry.timezone;
