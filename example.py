import singer
from signal_transformers.transform_runner import LOGGER, BaseSingerMessageTransformer
from singer import utils

FULL_TABLE = "FULL_TABLE"


class AddFullRefreshLoadDateSingerTransformer(BaseSingerMessageTransformer):
    def __init__(self, *args, **kwargs):
        self.full_table_load_at = utils.strftime(utils.now())
        super().__init__(*args, **kwargs)

        if not self.tap_config:
            LOGGER.warning(
                "No tap config provided to add-full-refresh-load-date transformer. No transformations applied"
            )

    def _do_schema_transform(self, message: singer.SchemaMessage) -> list[singer.SchemaMessage]:
        """Adds full_table_load_at to Singer schemas"""
        message.schema["properties"]["_sdc_full_table_load_at"] = {
            "type": ["null", "string"],
            "format": "date-time",
        }
        return [message]

    def _do_record_transform(self, message: singer.RecordMessage) -> list[singer.RecordMessage]:
        """Adds full_table_load_at to Singer records"""
        message.record["_sdc_full_table_load_at"] = self.full_table_load_at
        return [message]

    def transform(self, message: singer.Message) -> list[singer.Message]:
        if (
            self.tap_config
            and self.tap_config.get("default_replication_method") == FULL_TABLE
            # if default_replication_method is not defined, assume we're conditionally adding the transformer based on the run
            or not self.tap_config.get("default_replication_method")
        ):
            if isinstance(message, singer.RecordMessage):
                return self._do_record_transform(message)
            elif isinstance(message, singer.SchemaMessage):
                return self._do_schema_transform(message)

        return [message]
