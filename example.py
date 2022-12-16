def process_records(stream, mdata, max_modified, records, filter_field, fks):
    schema = stream.schema.to_dict()
    with metrics.record_counter(stream.tap_stream_id) as counter:
        for record in records:
            record_flat = {
                'id': record['id']
            }
            for prop, value in record['attributes'].items():
                if prop == 'id':
                    raise Exception(
                        'Error flattening Outeach record - conflict with `id` key')
                record_flat[prop] = value

            if 'relationships' in record:
                for prop, value in record['relationships'].items():
                    if 'data' not in value and 'links' not in value:
                        raise Exception(
                            'Only `data` or `links` expected in relationships')

                    fk_field_name = '{}Id'.format(prop)

                    if 'data' in value and fk_field_name in fks:
                        data_value = value['data']
                        if data_value is not None and 'id' not in data_value:
                            raise Exception(
                                'null or `id` field expected for `data` relationship')

                        if fk_field_name in record_flat:
                            raise Exception(
                                '`{}` exists as both an attribute and generated relationship name'.format(fk_field_name))

                        if data_value == None:
                            record_flat[fk_field_name] = None
                        else:
                            record_flat[fk_field_name] = data_value['id']