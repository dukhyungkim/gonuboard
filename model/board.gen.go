// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameBoard = "board"

// Board mapped from table <board>
type Board struct {
	BoTable               *string `gorm:"column:bo_table;type:VARCHAR(20)" json:"bo_table"`
	GrID                  *string `gorm:"column:gr_id;type:VARCHAR(255)" json:"gr_id"`
	BoSubject             *string `gorm:"column:bo_subject;type:VARCHAR(255)" json:"bo_subject"`
	BoMobileSubject       *string `gorm:"column:bo_mobile_subject;type:VARCHAR(255)" json:"bo_mobile_subject"`
	BoDevice              *string `gorm:"column:bo_device;type:VARCHAR(6)" json:"bo_device"`
	BoAdmin               *string `gorm:"column:bo_admin;type:VARCHAR(255)" json:"bo_admin"`
	BoListLevel           *int32  `gorm:"column:bo_list_level;type:INTEGER" json:"bo_list_level"`
	BoReadLevel           *int32  `gorm:"column:bo_read_level;type:INTEGER" json:"bo_read_level"`
	BoWriteLevel          *int32  `gorm:"column:bo_write_level;type:INTEGER" json:"bo_write_level"`
	BoReplyLevel          *int32  `gorm:"column:bo_reply_level;type:INTEGER" json:"bo_reply_level"`
	BoCommentLevel        *int32  `gorm:"column:bo_comment_level;type:INTEGER" json:"bo_comment_level"`
	BoUploadLevel         *int32  `gorm:"column:bo_upload_level;type:INTEGER" json:"bo_upload_level"`
	BoDownloadLevel       *int32  `gorm:"column:bo_download_level;type:INTEGER" json:"bo_download_level"`
	BoHTMLLevel           *int32  `gorm:"column:bo_html_level;type:INTEGER" json:"bo_html_level"`
	BoLinkLevel           *int32  `gorm:"column:bo_link_level;type:INTEGER" json:"bo_link_level"`
	BoCountDelete         *int32  `gorm:"column:bo_count_delete;type:INTEGER" json:"bo_count_delete"`
	BoCountModify         *int32  `gorm:"column:bo_count_modify;type:INTEGER" json:"bo_count_modify"`
	BoReadPoint           *int32  `gorm:"column:bo_read_point;type:INTEGER" json:"bo_read_point"`
	BoWritePoint          *int32  `gorm:"column:bo_write_point;type:INTEGER" json:"bo_write_point"`
	BoCommentPoint        *int32  `gorm:"column:bo_comment_point;type:INTEGER" json:"bo_comment_point"`
	BoDownloadPoint       *int32  `gorm:"column:bo_download_point;type:INTEGER" json:"bo_download_point"`
	BoUseCategory         *int32  `gorm:"column:bo_use_category;type:INTEGER" json:"bo_use_category"`
	BoCategoryList        *string `gorm:"column:bo_category_list;type:TEXT" json:"bo_category_list"`
	BoUseSideview         *int32  `gorm:"column:bo_use_sideview;type:INTEGER" json:"bo_use_sideview"`
	BoUseFileContent      *int32  `gorm:"column:bo_use_file_content;type:INTEGER" json:"bo_use_file_content"`
	BoUseSecret           *int32  `gorm:"column:bo_use_secret;type:INTEGER" json:"bo_use_secret"`
	BoUseDhtmlEditor      *int32  `gorm:"column:bo_use_dhtml_editor;type:INTEGER" json:"bo_use_dhtml_editor"`
	BoSelectEditor        *string `gorm:"column:bo_select_editor;type:VARCHAR(50)" json:"bo_select_editor"`
	BoUseRssView          *int32  `gorm:"column:bo_use_rss_view;type:INTEGER" json:"bo_use_rss_view"`
	BoUseGood             *int32  `gorm:"column:bo_use_good;type:INTEGER" json:"bo_use_good"`
	BoUseNogood           *int32  `gorm:"column:bo_use_nogood;type:INTEGER" json:"bo_use_nogood"`
	BoUseName             *int32  `gorm:"column:bo_use_name;type:INTEGER" json:"bo_use_name"`
	BoUseSignature        *int32  `gorm:"column:bo_use_signature;type:INTEGER" json:"bo_use_signature"`
	BoUseIPView           *int32  `gorm:"column:bo_use_ip_view;type:INTEGER" json:"bo_use_ip_view"`
	BoUseListView         *int32  `gorm:"column:bo_use_list_view;type:INTEGER" json:"bo_use_list_view"`
	BoUseListFile         *int32  `gorm:"column:bo_use_list_file;type:INTEGER" json:"bo_use_list_file"`
	BoUseListContent      *int32  `gorm:"column:bo_use_list_content;type:INTEGER" json:"bo_use_list_content"`
	BoTableWidth          *int32  `gorm:"column:bo_table_width;type:INTEGER" json:"bo_table_width"`
	BoSubjectLen          *int32  `gorm:"column:bo_subject_len;type:INTEGER" json:"bo_subject_len"`
	BoMobileSubjectLen    *int32  `gorm:"column:bo_mobile_subject_len;type:INTEGER" json:"bo_mobile_subject_len"`
	BoPageRows            *int32  `gorm:"column:bo_page_rows;type:INTEGER" json:"bo_page_rows"`
	BoMobilePageRows      *int32  `gorm:"column:bo_mobile_page_rows;type:INTEGER" json:"bo_mobile_page_rows"`
	BoNew                 *int32  `gorm:"column:bo_new;type:INTEGER" json:"bo_new"`
	BoHot                 *int32  `gorm:"column:bo_hot;type:INTEGER" json:"bo_hot"`
	BoImageWidth          *int32  `gorm:"column:bo_image_width;type:INTEGER" json:"bo_image_width"`
	BoSkin                *string `gorm:"column:bo_skin;type:VARCHAR(255)" json:"bo_skin"`
	BoMobileSkin          *string `gorm:"column:bo_mobile_skin;type:VARCHAR(255)" json:"bo_mobile_skin"`
	BoIncludeHead         *string `gorm:"column:bo_include_head;type:VARCHAR(255)" json:"bo_include_head"`
	BoIncludeTail         *string `gorm:"column:bo_include_tail;type:VARCHAR(255)" json:"bo_include_tail"`
	BoContentHead         *string `gorm:"column:bo_content_head;type:TEXT" json:"bo_content_head"`
	BoMobileContentHead   *string `gorm:"column:bo_mobile_content_head;type:TEXT" json:"bo_mobile_content_head"`
	BoContentTail         *string `gorm:"column:bo_content_tail;type:TEXT" json:"bo_content_tail"`
	BoMobileContentTail   *string `gorm:"column:bo_mobile_content_tail;type:TEXT" json:"bo_mobile_content_tail"`
	BoInsertContent       *string `gorm:"column:bo_insert_content;type:TEXT" json:"bo_insert_content"`
	BoGalleryCols         *int32  `gorm:"column:bo_gallery_cols;type:INTEGER" json:"bo_gallery_cols"`
	BoGalleryWidth        *int32  `gorm:"column:bo_gallery_width;type:INTEGER" json:"bo_gallery_width"`
	BoGalleryHeight       *int32  `gorm:"column:bo_gallery_height;type:INTEGER" json:"bo_gallery_height"`
	BoMobileGalleryWidth  *int32  `gorm:"column:bo_mobile_gallery_width;type:INTEGER" json:"bo_mobile_gallery_width"`
	BoMobileGalleryHeight *int32  `gorm:"column:bo_mobile_gallery_height;type:INTEGER" json:"bo_mobile_gallery_height"`
	BoUploadSize          *int32  `gorm:"column:bo_upload_size;type:INTEGER" json:"bo_upload_size"`
	BoReplyOrder          *int32  `gorm:"column:bo_reply_order;type:INTEGER" json:"bo_reply_order"`
	BoUseSearch           *int32  `gorm:"column:bo_use_search;type:INTEGER" json:"bo_use_search"`
	BoOrder               *int32  `gorm:"column:bo_order;type:INTEGER" json:"bo_order"`
	BoCountWrite          *int32  `gorm:"column:bo_count_write;type:INTEGER" json:"bo_count_write"`
	BoCountComment        *int32  `gorm:"column:bo_count_comment;type:INTEGER" json:"bo_count_comment"`
	BoWriteMin            *int32  `gorm:"column:bo_write_min;type:INTEGER" json:"bo_write_min"`
	BoWriteMax            *int32  `gorm:"column:bo_write_max;type:INTEGER" json:"bo_write_max"`
	BoCommentMin          *int32  `gorm:"column:bo_comment_min;type:INTEGER" json:"bo_comment_min"`
	BoCommentMax          *int32  `gorm:"column:bo_comment_max;type:INTEGER" json:"bo_comment_max"`
	BoNotice              *string `gorm:"column:bo_notice;type:TEXT" json:"bo_notice"`
	BoUploadCount         *int32  `gorm:"column:bo_upload_count;type:INTEGER" json:"bo_upload_count"`
	BoUseEmail            *int32  `gorm:"column:bo_use_email;type:INTEGER" json:"bo_use_email"`
	BoUseCert             *string `gorm:"column:bo_use_cert;type:VARCHAR(8)" json:"bo_use_cert"`
	BoUseSns              *int32  `gorm:"column:bo_use_sns;type:INTEGER" json:"bo_use_sns"`
	BoUseCaptcha          *int32  `gorm:"column:bo_use_captcha;type:INTEGER" json:"bo_use_captcha"`
	BoSortField           *string `gorm:"column:bo_sort_field;type:VARCHAR(255)" json:"bo_sort_field"`
	Bo1Subj               *string `gorm:"column:bo_1_subj;type:VARCHAR(255)" json:"bo_1_subj"`
	Bo2Subj               *string `gorm:"column:bo_2_subj;type:VARCHAR(255)" json:"bo_2_subj"`
	Bo3Subj               *string `gorm:"column:bo_3_subj;type:VARCHAR(255)" json:"bo_3_subj"`
	Bo4Subj               *string `gorm:"column:bo_4_subj;type:VARCHAR(255)" json:"bo_4_subj"`
	Bo5Subj               *string `gorm:"column:bo_5_subj;type:VARCHAR(255)" json:"bo_5_subj"`
	Bo6Subj               *string `gorm:"column:bo_6_subj;type:VARCHAR(255)" json:"bo_6_subj"`
	Bo7Subj               *string `gorm:"column:bo_7_subj;type:VARCHAR(255)" json:"bo_7_subj"`
	Bo8Subj               *string `gorm:"column:bo_8_subj;type:VARCHAR(255)" json:"bo_8_subj"`
	Bo9Subj               *string `gorm:"column:bo_9_subj;type:VARCHAR(255)" json:"bo_9_subj"`
	Bo10Subj              *string `gorm:"column:bo_10_subj;type:VARCHAR(255)" json:"bo_10_subj"`
	Bo1                   *string `gorm:"column:bo_1;type:VARCHAR(255)" json:"bo_1"`
	Bo2                   *string `gorm:"column:bo_2;type:VARCHAR(255)" json:"bo_2"`
	Bo3                   *string `gorm:"column:bo_3;type:VARCHAR(255)" json:"bo_3"`
	Bo4                   *string `gorm:"column:bo_4;type:VARCHAR(255)" json:"bo_4"`
	Bo5                   *string `gorm:"column:bo_5;type:VARCHAR(255)" json:"bo_5"`
	Bo6                   *string `gorm:"column:bo_6;type:VARCHAR(255)" json:"bo_6"`
	Bo7                   *string `gorm:"column:bo_7;type:VARCHAR(255)" json:"bo_7"`
	Bo8                   *string `gorm:"column:bo_8;type:VARCHAR(255)" json:"bo_8"`
	Bo9                   *string `gorm:"column:bo_9;type:VARCHAR(255)" json:"bo_9"`
	Bo10                  *string `gorm:"column:bo_10;type:VARCHAR(255)" json:"bo_10"`
}

// TableName Board's table name
func (*Board) TableName() string {
	return Prefix + TableNameBoard
}